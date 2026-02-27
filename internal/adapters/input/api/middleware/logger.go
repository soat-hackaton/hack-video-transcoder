package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/philipphahmann/hack-video-transcoder/pkg/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Extrai task_id se previamente injetado
		taskID := logger.GetCorrelationID(c.Request.Context())

		c.Next()

		latency := time.Since(start).Milliseconds()
		statusCode := c.Writer.Status()

		level := "INFO"
		if statusCode >= 400 && statusCode < 500 {
			level = "WARN"
		} else if statusCode >= 500 {
			level = "ERROR"
		}

		var errorMessage string
		if len(c.Errors) > 0 {
			errorMessage = c.Errors.ByType(gin.ErrorTypeAny).String()
		}

		attrs := []slog.Attr{
			slog.String("level", level),
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status_code", statusCode),
			slog.Int64("latency_ms", latency),
		}

		if taskID != "" {
			attrs = append(attrs, slog.String("task_id", taskID))
		}
		if errorMessage != "" {
			attrs = append(attrs, slog.String("error_message", errorMessage))
		}

		// Log estructurado de acordo com Golden Signals
		if level == "ERROR" {
			slog.LogAttrs(c.Request.Context(), slog.LevelError, "HTTP request failed", attrs...)
		} else if level == "WARN" {
			slog.LogAttrs(c.Request.Context(), slog.LevelWarn, "HTTP request warning", attrs...)
		} else {
			slog.LogAttrs(c.Request.Context(), slog.LevelInfo, "HTTP request completed", attrs...)
		}
	}
}
