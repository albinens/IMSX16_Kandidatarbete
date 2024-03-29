package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

func IsProduction() bool {
	return os.Getenv("ENVIRONMENT") == "production"
}

func LogFatal(s string, v ...any) {
	slog.Error(s, v...)
	os.Exit(1)
}

type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func WriteHttpError(w http.ResponseWriter, message string, statusCode int) {
	data, _ := json.Marshal(errorResponse{Message: message, Code: statusCode})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(data)
	if err != nil {
		slog.Error("Failed to write error response: ", err)
	}
}
