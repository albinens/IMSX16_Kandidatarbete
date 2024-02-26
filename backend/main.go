package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("GET /", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
