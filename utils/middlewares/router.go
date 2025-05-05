package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

// GinZapLogger - GinZapLogger is a middleware that logs HTTP requests using zap logger.
func GinZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		logger.Info("HTTP request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}

func TracingMiddleware() gin.HandlerFunc {
	tracer := otel.Tracer("gin-service")

	return func(c *gin.Context) {
		// Create a new span for the incoming HTTP request
		ctx, span := tracer.Start(c.Request.Context(), c.Request.Method+" "+c.Request.URL.Path)
		defer span.End()

		// Inject the context into the request
		c.Request = c.Request.WithContext(ctx)

		// Process the request
		c.Next()
	}
}
