package logger

import (
	"context"
	"log/slog"
	"os"
)

type correlationIDKey struct{}

// WithCorrelationID injects the task_id into the context.
func WithCorrelationID(ctx context.Context, taskID string) context.Context {
	return context.WithValue(ctx, correlationIDKey{}, taskID)
}

// GetCorrelationID extracts the task_id from the context.
func GetCorrelationID(ctx context.Context) string {
	if val, ok := ctx.Value(correlationIDKey{}).(string); ok {
		return val
	}
	return ""
}

// Setup initializes the global slog instance with JSON formatting.
func Setup() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// Correlate logger wraps the base handler to automatically append task_id from context
	logger := slog.New(&correlationHandler{Handler: handler})
	slog.SetDefault(logger)
}

type correlationHandler struct {
	slog.Handler
}

// Handle adds the task_id to the log record if it exists in the context.
func (h *correlationHandler) Handle(ctx context.Context, r slog.Record) error {
	taskID := GetCorrelationID(ctx)
	if taskID != "" {
		r.AddAttrs(slog.String("task_id", taskID))
	}
	return h.Handler.Handle(ctx, r)
}
