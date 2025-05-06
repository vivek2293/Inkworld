package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vivek2293/Inkworld/constants"
	"github.com/vivek2293/Inkworld/utils/monitoring"
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
	tracer := otel.Tracer(constants.TracerName)

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

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // process request

		duration := time.Since(start).Seconds()
		path := c.FullPath() // avoids high cardinality from query strings
		if path == "" {
			path = c.Request.URL.Path // fallback
		}

		monitoring.HttpRequestsTotal.WithLabelValues(c.Request.Method, path, http.StatusText(c.Writer.Status())).Inc()
		monitoring.HttpRequestDuration.WithLabelValues(c.Request.Method, path).Observe(duration)
	}
}
