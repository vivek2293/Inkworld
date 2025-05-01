package logger

import "go.uber.org/zap"

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

func Sync() error {
	return zap.L().Sync()
}
