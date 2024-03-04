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
	mux.HandleFunc("POST /api/add-room", addRoomHandler)

	wrappedMux := logger.NewRequestLoggerMiddleware(mux)

	slog.Info("Starting server with", "port", env.Port)
	utils.LogFatal("Server crashed with error: ", http.ListenAndServe(":"+env.Port, wrappedMux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	rooms, err := room.StatusOfAllRooms()
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to determine room status", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(rooms)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to convert room status to JSON", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func addRoomHandler(w http.ResponseWriter, r *http.Request) {
	var roomToAdd struct {
		Name string `json:"name"`
		Sensor string `json:"sensor"`
		Building string `json:"building"`
	} 
	if err := json.NewDecoder(r.Body).Decode(&roomToAdd); err != nil {
		slog.ErrorContext(r.Context(), "Failed to decode request body", "error", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := room.AddRoom(roomToAdd.Name, roomToAdd.Sensor, roomToAdd.Building); err != nil {
		slog.ErrorContext(r.Context(), "Failed to add room", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
