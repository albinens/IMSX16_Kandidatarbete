package api

import (
	"log/slog"
	"net/http"

	"example.com/m/v2/room"
	"example.com/m/v2/utils"
)

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if err := room.DeleteRoom(name); err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to delete room", "error", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
