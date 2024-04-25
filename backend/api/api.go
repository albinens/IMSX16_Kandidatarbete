package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"example.com/m/v2/auth"
	"example.com/m/v2/env"
	"example.com/m/v2/logger"
	"example.com/m/v2/utils"
)

func Init() {
	mux := http.NewServeMux()

	apiDocs := http.StripPrefix("/api", http.FileServer(http.Dir("../docs")))
	mux.Handle("GET /", apiDocs)

	mux.HandleFunc("GET /api/current", currentStatus)
	mux.HandleFunc("GET /api/current/{room}", roomStatus)
	mux.HandleFunc("GET /api/stats/daily-average/{from}/{to}", dailyAverage)
	mux.HandleFunc("GET /api/rooms", allRooms)

	mux.Handle("POST /api/add-room", auth.TokenAuthMiddlewareFunc(addRoom))
	mux.HandleFunc("POST /api/report/status", status)
	mux.HandleFunc("POST /api/auth/setup", setupAuth)
	mux.Handle("POST /api/auth/key/create", auth.TokenAuthMiddlewareFunc(createKey))

	mux.Handle("DELETE /api/remove-room/{name}", auth.TokenAuthMiddlewareFunc(deleteRoom))

	wrappedMux := logger.NewRequestLoggerMiddleware(mux)

	slog.Info("Starting server", "port", env.Port)
	utils.LogFatal("Server crashed with error: ", http.ListenAndServe(":"+env.Port, wrappedMux))
}

func sendJSONResponse(w http.ResponseWriter, r *http.Request, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to create JSON response", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
