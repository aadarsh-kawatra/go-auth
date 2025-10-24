package auth

import (
	"encoding/json"
	"net/http"

	"crud/pkg/utils"
)

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

	validationErr := utils.ValidateStruct(req)
	if len(validationErr) > 0 {
		utils.HandleValidationError(w, validationErr)
		return
	}

	token, authErr := AuthenticateUserService(req.Email, req.Password)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := LoginResponse{}
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

	validationErr := utils.ValidateStruct(req)
	if len(validationErr) > 0 {
		utils.HandleValidationError(w, validationErr)
		return
	}

	token, authErr := RegisterUserService(req.FirstName, req.LastName, req.Email, req.Password)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := RegisterResponse{}
	resp.Message = "Registered"
	resp.Token = token

	json.NewEncoder(w).Encode(resp)
}
