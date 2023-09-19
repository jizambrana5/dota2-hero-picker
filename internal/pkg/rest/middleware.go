package rest

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const dateFormat = "2006-01-02T15:04:05"

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds.",
		},
		[]string{"method", "path"},
	)
)

func setupLogger() (*zap.Logger, error) {
	// Check the environment (e.g., "development" or "production")
	env := os.Getenv("ENVIRONMENT")

	// Configure the logger differently based on the environment
	var logger *zap.Logger
	var err error

	if env == "development" {
		// Development environment configuration
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(dateFormat)
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.DebugLevel,
		)
		logger = zap.New(core, zap.AddCaller())
	} else {
		// Production environment configuration
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(dateFormat)
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.InfoLevel,
		)
		logger = zap.New(core)
	}

	if err != nil {
		panic("Failed to create logger: " + err.Error())
	}
	defer logger.Sync() // nolint

	logger.Info("Logger initialized", zap.String("environment", env))
	return logger, nil
}

func loggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start a timer to measure request processing time
		startTime := time.Now()

		// Generate a unique request ID
		requestID := uuid.New().String()

		// Log the incoming request with date and time
		logger.Info("Incoming Request",
			zap.String("request_id", requestID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.Time("timestamp", startTime), // Include date and time
		)

		// Store the request ID in the context for future use
		c.Set("request_id", requestID)

		// Continue processing the request
		c.Next()

		// Log the response and processing time
		logger.Info("Response",
			zap.String("request_id", requestID),
			zap.Int("status_code", c.Writer.Status()),
			zap.Duration("processing_time", time.Since(startTime)),
			zap.Time("timestamp", time.Now()), // Include date and time
		)
	}
}

func MetricsMiddleware() gin.HandlerFunc {
	// Register the counter with Prometheus
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process the request
		c.Next()

		// Calculate request duration
		duration := time.Since(startTime).Seconds()

		// Increment the total requests counter
		// Use the getStatusString function to get the string representation of the status code.
		httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath(), getStatusString(c.Writer.Status())).Inc()

		// Observe request duration
		httpRequestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration)
	}
}

// Define a function to map integer status codes to string representations.
func getStatusString(status int) string {
	switch status {
	case 200:
		return "200"
	case 201:
		return "201"
	case 400:
		return "400"
	case 404:
		return "404"
	case 500:
		return "500"
	default:
		return "other"
	}
}
