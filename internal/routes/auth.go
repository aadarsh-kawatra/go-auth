package routes

import (
	"net/http"

	"crud/internal/auth"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/login", auth.LoginHandler)
}
