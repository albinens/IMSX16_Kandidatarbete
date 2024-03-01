package logger

import (
	"context"
	"log/slog"
	"os"

	"example.com/m/v2/utils"
	"github.com/google/uuid"
)

type contextKey int

const (
	TraceCtxKey contextKey = iota + 1
)

type RequestLogger struct {
	slog.Handler
}

func Configure() {
	var handler slog.Handler
	if !utils.IsProduction() {
		handler = slog.NewTextHandler(os.Stdout, nil)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}
	slog.SetDefault(slog.New(RequestLogger{handler}))
}

func GenerateTraceID() string {
	return uuid.New().String()
}

func (h RequestLogger) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(TraceCtxKey).(string); ok {
		r.Add("RequestID", slog.StringValue(traceID))
	}

	return h.Handler.Handle(ctx, r)
}
