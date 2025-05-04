package logger

import (
	"go.uber.org/zap"
)

// Debug - Debug logs a message at Debug level. Disabled in production.
func Debug(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

// Info - Info logs a message at Info level. Default level priority.
func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

// Warn - Warn logs a message at Warn level. Non-fatal errors.
func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

// Error - Error logs a message at Error level. Required for error handling.
func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

// DPanic - DPanic logs a message at DPanic level. It panics after logging. Usually used in development.
func DPanic(msg string, fields ...zap.Field) {
	zap.L().DPanic(msg, fields...)
}

// Panic - Panic logs a message at Panic level. It panics after logging.
func Panic(msg string, fields ...zap.Field) {
	zap.L().Panic(msg, fields...)
}

// Fatal - Fatal logs a message at Fatal level. It exits the program after logging.
func Fatal(msg string, fields ...zap.Field) {
	zap.L().Fatal(msg, fields...)
}

// Sync - Sync flushes any buffered log entries. It's a good practice to call this before the program exits.
func Sync() error {
	return zap.L().Sync()
}

// GetLogger - GetLogger returns the global logger instance. It's safe for concurrent use.
func GetLogger() *zap.Logger {
	return zap.L()
}
