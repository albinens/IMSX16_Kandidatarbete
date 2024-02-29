package main

import (
	"encoding/json"
	"net/http"

	"example.com/m/v2/database"
	"example.com/m/v2/env"
	"example.com/m/v2/room"
	"example.com/m/v2/seeder"
)

func main() {
	env.Load()
	database.InitSQL()
	database.InitTS()
	seeder.SeedDevelopmentData()

	http.HandleFunc("GET /", handler)
	http.HandleFunc("GET /api/current", currentHandler)

	room.StatusOfAllRooms()

	http.ListenAndServe(env.Port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	rooms := room.StatusOfAllRooms()
	data, err := json.Marshal(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
