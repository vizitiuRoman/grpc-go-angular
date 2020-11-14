package server

import (
	"github.com/user-service/pkg/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger() *zap.SugaredLogger {
	var config zap.Config
	if settings.GetEnvironment() == "prod" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger.Sugar()
}
