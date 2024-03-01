package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"example.com/m/v2/database"
	"example.com/m/v2/env"
	"example.com/m/v2/logger"
	"example.com/m/v2/room"
	"example.com/m/v2/seeder"
	"example.com/m/v2/utils"
)

func main() {
	env.Load()
	logger.Configure()
	if err := database.InitSQL(); err != nil {
		utils.LogFatal("Failed to initialize SQL: ", err)
	}
	database.InitTimeSeries()
	if err := seeder.SeedDevelopmentData(); err != nil {
		utils.LogFatal("Failed to seed development data: ", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler)
	mux.HandleFunc("GET /api/current", currentHandler)

	wrappedMux := logger.NewRequestLoggerMiddleware(mux)

	slog.Info("Starting server with", "port", env.Port)
	utils.LogFatal("Server crashed with error: ", http.ListenAndServe(":"+env.Port, wrappedMux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	rooms, err := room.StatusOfAllRooms(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
