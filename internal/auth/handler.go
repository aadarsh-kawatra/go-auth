package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"crud/pkg/utils"
)

func validateLoginPayload(body LoginRequest) error {
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
	var req LoginRequest

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

	token, authErr := AuthenticateUserService(req.Email, req.Password)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := LoginResponse{}
	resp.Code = http.StatusOK
	resp.Message = "Logged In"
	resp.Token = token

	json.NewEncoder(w).Encode(resp)
}

func RegisterHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req RegisterRequest

	decodeErr := json.NewDecoder(r.Body).Decode(&req)
	if decodeErr != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, authErr := RegisterUserService(req.FirstName, req.LastName, req.Email, req.Password)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := RegisterResponse{}
	resp.Code = http.StatusCreated
	resp.Message = "Registered"
	resp.Token = token

	json.NewEncoder(w).Encode(resp)
}
