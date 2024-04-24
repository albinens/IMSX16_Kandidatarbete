package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	g1gateway "example.com/m/v2/api/g1_gateway"
	"example.com/m/v2/room"
	"example.com/m/v2/sensor"
	"example.com/m/v2/utils"
)

func status(w http.ResponseWriter, r *http.Request) {
	data, err := g1gateway.ParseStatus(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, statusData := range data {
		statusRoom, err := sensor.RoomFromMac(statusData.MacAdress)
		if err != nil {
			utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
			slog.ErrorContext(r.Context(), "Failed to get room from mac", "error", err)
			return
		}

		if statusRoom == "" {
			slog.WarnContext(r.Context(), "Room not found", "room", statusData.MacAdress)
			continue
		}

		room.AddStatus(statusRoom, statusData.NrOfPeople)
	}

	w.WriteHeader(http.StatusCreated)
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

	if roomToAdd.Name == "" || roomToAdd.Sensor == "" || roomToAdd.Building == "" {
		utils.WriteHttpError(w, "No fields can be empty", http.StatusBadRequest)
		slog.DebugContext(r.Context(), "Empty fields sent to create room", "room", roomToAdd)
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
