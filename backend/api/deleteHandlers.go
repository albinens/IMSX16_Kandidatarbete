package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"example.com/m/v2/auth"
	"example.com/m/v2/room"
	"example.com/m/v2/utils"
)

func removeGatewayLogin(w http.ResponseWriter, r *http.Request) {
	var login struct {
		Username string `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		utils.WriteHttpError(w, "Request body is malformed", http.StatusBadRequest)
		slog.DebugContext(r.Context(), "Failed to decode request body", "error", err)
		return
	}

	if login.Username == "" {
		utils.WriteHttpError(w, "No fields can be empty", http.StatusBadRequest)
		slog.DebugContext(r.Context(), "Empty fields sent to remove gateway login", "login", login)
		return
	}

	if err := auth.DeleteGatewayUser(r.Context(), login.Username); err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to remove gateway login", "error", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteApiKey(w http.ResponseWriter, r *http.Request) {
	var key struct {
		Key string `json:"key"`
	}

	if err := json.NewDecoder(r.Body).Decode(&key); err != nil {
		utils.WriteHttpError(w, "Request body is malformed", http.StatusBadRequest)
		slog.DebugContext(r.Context(), "Failed to decode request body", "error", err)
		return
	}

	if err := auth.DeleteApiKey(r.Context(), key.Key); err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to delete api key", "error", err)
		return
	}

	w.WriteHeader(http.StatusOK)
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
