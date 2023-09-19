package logs

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger initializes the logger.
func InitLogger(environment string) {
	var cfg zap.Config

	if environment == "development" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		log.Panic("Failed to create logger: " + err.Error())
	}
	Logger = logger
}
