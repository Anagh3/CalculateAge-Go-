package logger

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment() // use zap.NewProduction() in prod
	return logger
}
