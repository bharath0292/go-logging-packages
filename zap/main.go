package main

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Basic
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()
	sugar.Info("Hello from zap logger")

	// Customize
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err = loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	sugar = logger.Sugar()
	sugar.Info("Hello from Sugar zap logger")
	sugar.Infow("Hello from zap logger",
		"tag", "hello_zap",
		"service", "logger",
	)

	l := sugar.Desugar()
	l.Info("Hello from Desugar zap logger",
		zap.String("tag", "hello_zap"),
		zap.Int("count", 10),
	)
}
