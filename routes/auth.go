package routes

import (
	"net/http"

	"crud/handlers"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/login", handlers.LoginHandler)
}
