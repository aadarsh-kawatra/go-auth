package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"crud/infrastructure/db"
	"crud/infrastructure/routes"
)

var PORT string = "8000"

type HealthResponse struct {
	Message string
}

func main() {
	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	routes.RegisterAuthRoutes(mux)
	routes.RegisterUserRoutes(mux)

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
