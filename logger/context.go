package logger

import (
	"context"
)

type contextKey string

const (
	contextLogger = contextKey("contextLogger")
)

func FromContext(ctx context.Context) *Logger {
	v := ctx.Value(contextLogger)
	if v == nil {
		return New()
	}

	return v.(*Logger)
}

func WithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, contextLogger, l)
}
