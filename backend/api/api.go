package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/m/v2/auth"
	"example.com/m/v2/env"
	"example.com/m/v2/logger"
	"example.com/m/v2/room"
	"example.com/m/v2/utils"
)

func Init() {
	mux := http.NewServeMux()

	apiDocs := http.StripPrefix("/api", http.FileServer(http.Dir("../docs")))
	mux.Handle("GET /api/", apiDocs)

	mux.HandleFunc("GET /api/current", currentStatus)
	mux.HandleFunc("GET /api/current/{room}", roomStatus)
	mux.HandleFunc("GET /api/stats/daily-average/{from}/{to}", dailyAverage)

	mux.Handle("POST /api/add-room", auth.TokenAuthMiddlewareFunc(addRoom))
	mux.Handle("DELETE /api/remove-room/{name}", auth.TokenAuthMiddlewareFunc(deleteRoom))

	wrappedMux := logger.NewRequestLoggerMiddleware(mux)

	slog.Info("Starting server", "port", env.Port)
	utils.LogFatal("Server crashed with error: ", http.ListenAndServe(":"+env.Port, wrappedMux))
}

func currentStatus(w http.ResponseWriter, r *http.Request) {
	rooms, err := room.StatusOfAllRooms()
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to determine room status", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, r, rooms)
}

func dailyAverage(w http.ResponseWriter, r *http.Request) {
	from := r.PathValue("from")
	to := r.PathValue("to")

	fromInt, err := strconv.Atoi(from)
	if err != nil {
		utils.WriteHttpError(w, "Invalid 'from' date", http.StatusBadRequest)
		slog.Debug("Invalid 'from' date", "from", from)
		return
	}

	toInt, err := strconv.Atoi(to)
	if err != nil {
		utils.WriteHttpError(w, "Invalid 'to' date", http.StatusBadRequest)
		slog.Debug("Invalid 'to' date", "to", to)
		return
	}

	fromTime := time.Unix(int64(fromInt), 0)
	toTime := time.Unix(int64(toInt), 0)

	if fromTime.After(toTime) {
		utils.WriteHttpError(w, "'from' date must be before 'to' date", http.StatusBadRequest)
		slog.Debug("Invalid date range", "from", from, "to", to)
		return
	}

	data, err := room.RoomOccupancyPerDayOfWeek(fromTime, toTime)
	if err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to determine daily averages", "error", err)
		return
	}

	type responseData struct {
		RoomName      string             `json:"roomName"`
		DailyAverages map[string]float32 `json:"dailyAverages"`
	}

	var convertedData []responseData
	for room, averages := range *data {
		convertedData = append(convertedData, responseData{
			RoomName:      room,
			DailyAverages: averages,
		})
	}

	if len(convertedData) == 0 {
		utils.WriteHttpError(w, "No data available for given date range", http.StatusNotFound)
		slog.Debug("No data available for given date range", "from", from, "to", to)
		return
	}

	sendJSONResponse(w, r, convertedData)
}

func addRoom(w http.ResponseWriter, r *http.Request) {
	var roomToAdd struct {
		Name     string `json:"name"`
		Sensor   string `json:"sensor"`
		Building string `json:"building"`
	}

	if err := json.NewDecoder(r.Body).Decode(&roomToAdd); err != nil {
		utils.WriteHttpError(w, "Request body is malformed", http.StatusBadRequest)
		slog.ErrorContext(r.Context(), "Failed to decode request body", "error", err)
		return
	}

	if err := room.AddRoom(roomToAdd.Name, roomToAdd.Sensor, roomToAdd.Building); err != nil {
		if strings.HasPrefix(err.Error(), "pq: duplicate key value violates unique constraint") {
			utils.WriteHttpError(w, "Room already exists", http.StatusConflict)
			slog.DebugContext(r.Context(), "Room already exists", "room", roomToAdd.Name)
			return
		}

		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to add room", "error", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if err := room.DeleteRoom(name); err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to delete room", "error", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func roomStatus(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("room")
	status, err := room.StatusOfRoom(name)
	if err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to determine room status", "error", err)
		return
	}

	sendJSONResponse(w, r, status)
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
