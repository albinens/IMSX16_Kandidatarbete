package main

import (
	"net/http"

	"example.com/m/v2/database"
	"example.com/m/v2/room"
)

func main() {
	database.InitSQL()
	database.InitTS()

	http.HandleFunc("GET /", handler)
	http.HandleFunc("GET /api/current", currentHandler)

	room.GetStatusOfAllRooms()

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func currentHandler(w http.ResponseWriter, r *http.Request) {

}
