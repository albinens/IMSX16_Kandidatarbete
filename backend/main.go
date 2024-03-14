package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"example.com/m/v2/auth"
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

	mux.HandleFunc("GET /api/current", currentHandler)
	mux.HandleFunc("GET /api/current/{room}", currentRoomHandler)

	mux.Handle("POST /api/add-room", auth.TokenAuthMiddleware(http.HandlerFunc(addRoomHandler)))
	mux.Handle("DELETE /api/remove-room/{name}", auth.TokenAuthMiddleware(http.HandlerFunc(deleteRoomHandler)))

	wrappedMux := logger.NewRequestLoggerMiddleware(mux)

	slog.Info("Starting server with", "port", env.Port)
	utils.LogFatal("Server crashed with error: ", http.ListenAndServe(":"+env.Port, wrappedMux))
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
		Name     string `json:"name"`
		Sensor   string `json:"sensor"`
		Building string `json:"building"`
	}
	if err := json.NewDecoder(r.Body).Decode(&roomToAdd); err != nil {
		slog.ErrorContext(r.Context(), "Failed to decode request body", "error", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := room.AddRoom(roomToAdd.Name, roomToAdd.Sensor, roomToAdd.Building); err != nil {
		if strings.HasPrefix(err.Error(), "pq: duplicate key value violates unique constraint") {
			http.Error(w, "Room already exists", http.StatusConflict)
			slog.DebugContext(r.Context(), "Room already exists", "room", roomToAdd.Name)
			return
		}

		slog.ErrorContext(r.Context(), "Failed to add room", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if err := room.DeleteRoom(name); err != nil {
		slog.ErrorContext(r.Context(), "Failed to delete room", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func currentRoomHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("room")
	status, err := room.StatusOfRoom(name)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to determine room status", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(status)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to convert room status to JSON", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
