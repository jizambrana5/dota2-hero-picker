package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLoggingMiddleware(t *testing.T) {
	// Create a temporary Gin router for testing
	router := gin.New()

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Create a dummy request
	req, _ := http.NewRequest("GET", "/test", nil)

	// Create a buffer to capture the log output
	var logOutput []string
	writeSyncer := zapcore.AddSync(&logSliceWriter{&logOutput})

	// Create a test logger that writes to the buffer
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writeSyncer,
		zap.NewAtomicLevelAt(zap.DebugLevel),
	))

	// Add the logging middleware to the router
	router.Use(loggingMiddleware(logger))

	// Define a test route that uses the middleware
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test Response")
	})

	// Perform a test request
	router.ServeHTTP(w, req)

	// Check the recorded log entries
	if len(logOutput) < 2 {
		t.Fatal("Expected at least 2 log entries, got", len(logOutput))
	}

	// Check for "Incoming Request" and "Response" log entries in JSON format
	var logEntry1 map[string]interface{}
	if err := json.Unmarshal([]byte(logOutput[0]), &logEntry1); err != nil {
		t.Fatal("Failed to unmarshal log entry 1:", err)
	}

	if logEntry1["msg"] != "Incoming Request" {
		t.Error("Expected 'Incoming Request' in log entry 1, got", logEntry1["msg"])
	}

	var logEntry2 map[string]interface{}
	if err := json.Unmarshal([]byte(logOutput[1]), &logEntry2); err != nil {
		t.Fatal("Failed to unmarshal log entry 2:", err)
	}

	if logEntry2["msg"] != "Response" {
		t.Error("Expected 'Response' in log entry 2, got", logEntry2["msg"])
	}
}

// Custom writer to capture log entries
type logSliceWriter struct {
	logs *[]string
}

func (w *logSliceWriter) Write(p []byte) (n int, err error) {
	*w.logs = append(*w.logs, string(p))
	return len(p), nil
}
