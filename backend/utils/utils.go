package utils

import (
	"log/slog"
	"os"
)

func IsProduction() bool {
	return os.Getenv("ENVIRONMENT") == "production"
}

func LogFatal(s string, v ...any) {
	slog.Error(s, v...)
	os.Exit(1)
}
