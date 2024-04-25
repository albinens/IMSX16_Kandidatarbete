package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func IsProduction() bool {
	return os.Getenv("ENVIRONMENT") == "production"
}

func LogFatal(s string, v ...any) {
	slog.Error(s, v...)
	os.Exit(1)
}

func ValidTimeUnit(unit string) bool {
	_, err := strconv.Atoi(unit[:len(unit)-1])
	if err != nil {
		return false
	}

	switch unit[len(unit)-1] {
	case 's', 'm', 'h', 'd':
		return true
	default:
		return false
	}
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

func IsInList[T comparable](list []T, item T) bool {
	for _, listItem := range list {
		if listItem == item {
			return true
		}
	}
	return false
}

func IsInListFunc[T any](list []T, f func(T) bool) bool {
	for _, listItem := range list {
		if f(listItem) {
			return true
		}
	}
	return false
}
