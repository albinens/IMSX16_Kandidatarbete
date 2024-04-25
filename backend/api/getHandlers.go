package api

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"example.com/m/v2/room"
	"example.com/m/v2/utils"
)

func rawSerialData(w http.ResponseWriter, r *http.Request) {
	from := r.PathValue("from")
	to := r.PathValue("to")
	resolution := r.PathValue("resolution")

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

	if !utils.ValidTimeUnit(resolution) {
		utils.WriteHttpError(w, "Invalid resolution", http.StatusBadRequest)
		slog.Debug("Invalid resolution", "resolution", resolution)
		return
	}

	result, err := room.RawDataBetween(fromTime, toTime, resolution)
	if err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to determine raw serial data", "error", err)
		return
	}

	type responseDataType struct {
		Timestamp int64   `json:"timestamp"`
		Occupancy float64 `json:"occupancy"`
	}

	type responseType struct {
		RoomName string             `json:"roomName"`
		Data     []responseDataType `json:"data"`
	}

	response := make([]responseType, 0)

	for room, data := range result {
		convertedData := make([]responseDataType, 0, len(data))
		for _, entry := range data {
			convertedData = append(convertedData, responseDataType{
				Timestamp: entry.Timestamp.Unix(),
				Occupancy: entry.Occupancy,
			})
		}

		response = append(response, responseType{RoomName: room,
			Data: convertedData,
		})
	}

	sendJSONResponse(w, r, response)
}

func allRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := room.AllRooms()
	if err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to get all rooms", "error", err)
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

func currentStatus(w http.ResponseWriter, r *http.Request) {
	rooms, err := room.StatusOfAllRooms()
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to determine room status", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, r, rooms)
}
