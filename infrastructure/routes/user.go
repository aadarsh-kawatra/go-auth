package routes

import (
	"net/http"

	"crud/infrastructure/middlewares"
	"crud/internal/user"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	mux.Handle("GET /user/", middlewares.AuthMiddleware(http.HandlerFunc(user.GetProfileHandler)))
}
