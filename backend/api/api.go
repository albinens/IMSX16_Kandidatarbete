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

	// apiDocs := http.StripPrefix("/api", http.FileServer(http.Dir("../docs")))
	// mux.Handle("GET /api", apiDocs)

	frontend := http.StripPrefix("/", http.FileServer(http.Dir("./public")))
	listRooms := http.StripPrefix("/listRooms", http.FileServer(http.Dir("./public")))
	listrooms := http.StripPrefix("/listrooms", http.FileServer(http.Dir("./public")))
	dataBoard := http.StripPrefix("/dataBoard", http.FileServer(http.Dir("./public")))
	sensors := http.StripPrefix("/sensors", http.FileServer(http.Dir("./public")))
	about := http.StripPrefix("/about", http.FileServer(http.Dir("./public")))

	mux.Handle("GET /", frontend)
	mux.Handle("GET /listRooms", listRooms)
	mux.Handle("GET /listrooms", listrooms)
	mux.Handle("GET /dataBoard", dataBoard)
	mux.Handle("GET /sensors", sensors)
	mux.Handle("GET /about", about)

	mux.HandleFunc("GET /api/current", currentStatus)
	mux.HandleFunc("GET /api/current/{room}", roomStatus)
	mux.HandleFunc("GET /api/stats/daily-average/{from}/{to}", dailyAverage)
	mux.Handle("GET /api/stats/raw-serial/{from}/{to}/{resolution}", auth.TokenAuthMiddlewareFunc(rawSerialData))
	mux.HandleFunc("GET /api/rooms", allRooms)

	mux.Handle("POST /api/add-room", auth.TokenAuthMiddlewareFunc(addRoom))
	mux.HandleFunc("POST /api/report/status", status)
	mux.HandleFunc("POST /api/auth/setup", setupAuth)
	mux.Handle("POST /api/auth/key/create", auth.TokenAuthMiddlewareFunc(createKey))
	mux.Handle("POST /api/auth/gateway/create", auth.TokenAuthMiddlewareFunc(createGatewayLogin))
	mux.Handle("POST /api/auth/gateway/remove", auth.TokenAuthMiddlewareFunc(removeGatewayLogin))

	mux.Handle("DELETE /api/remove-room/{name}", auth.TokenAuthMiddlewareFunc(deleteRoom))
	mux.Handle("DELETE /api/auth/key/revoke", auth.TokenAuthMiddlewareFunc(deleteApiKey))

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
