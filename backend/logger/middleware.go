package logger

import (
	"context"
	// "log/slog"
	"net/http"
	// "time"
)

type RequestLoggerMiddleware struct {
	handler http.Handler
}

func (l *RequestLoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	traceID := GenerateTraceID()
	ctx = context.WithValue(ctx, TraceCtxKey, traceID)
	r = r.WithContext(ctx)

	// slog.InfoContext(ctx, "Request started", "method", r.Method, "path", r.URL.Path)
	// start := time.Now()
	l.handler.ServeHTTP(w, r)
	// slog.InfoContext(ctx, "Request finished", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
}

func NewRequestLoggerMiddleware(handlerToWrap http.Handler) *RequestLoggerMiddleware {
	return &RequestLoggerMiddleware{handlerToWrap}
}
