package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/m/v2/database"
	"example.com/m/v2/env"
	"example.com/m/v2/room"
	"example.com/m/v2/seeder"
)

func main() {
	env.Load()
	if err := database.InitSQL(); err != nil {
		fmt.Printf("Failed to initialize SQL: %s\n", err.Error())
	}
	database.InitTimeSeries()
	if err := seeder.SeedDevelopmentData(); err != nil {
		fmt.Printf("Failed to seed development data: %s\n", err.Error())
	}

	http.HandleFunc("GET /", handler)
	http.HandleFunc("GET /api/current", currentHandler)

	http.ListenAndServe(env.Port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	rooms, err := room.StatusOfAllRooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
