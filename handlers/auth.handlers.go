package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"crud/models"
	"crud/services"
	"crud/utils"
)

func validateLoginPayload(body models.LoginRequest) error {
	if body.Email == "" || body.Password == "" {
		return errors.New("email & password required")
	}
	if !utils.IsEmail(body.Email) {
		return errors.New("invalid email")
	}
	return nil
}

func LoginHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req models.LoginRequest

	decodeErr := json.NewDecoder(r.Body).Decode(&req)
	if decodeErr != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	validationErr := validateLoginPayload(req)
	if validationErr != nil {
		http.Error(w, validationErr.Error(), http.StatusBadRequest)
		return
	}

	token, authErr := services.AuthenticateUserService(req.Email, req.Password)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := models.LoginResponse{
		Code:    http.StatusOK,
		Message: "Logged In",
		Token:   token,
	}
	json.NewEncoder(w).Encode(resp)
}
