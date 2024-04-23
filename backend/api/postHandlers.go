package api

import (
	"encoding/hex"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"example.com/m/v2/room"
	"example.com/m/v2/utils"
)

var global []int = make([]int, 0)

func status(w http.ResponseWriter, r *http.Request) {
	data := make([]map[string]interface{}, 0)
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.WriteHttpError(w, "Request body is malformed", http.StatusBadRequest)
		slog.ErrorContext(r.Context(), "Failed to decode request body", "error", err)
		return
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	for _, dataItem := range data {
		rawData, ok := dataItem["rawData"]
		if !ok {
			continue
		}

		stringData := rawData.(string)
		if rawData == "" {
			continue
		}

		bytes, err := hex.DecodeString(stringData)
		if err != nil {
			panic(err)
		}

		nr := int((bytes[11] << 4) >> 4)
		global = append(global, nr)
	}
}

func sendJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(global)
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
