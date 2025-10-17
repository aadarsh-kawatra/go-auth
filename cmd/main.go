package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var PORT string = "8000"

type HealthResponse struct {
	Message string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	log.Println("Server running on http://localhost:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}

func healthHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	resp := HealthResponse{Message: "Hello World"}

	json.NewEncoder(w).Encode(resp)
}
