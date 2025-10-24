package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseStruct struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func HandleErrorResponse(w http.ResponseWriter, code int, message string, errors []string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := ErrorResponseStruct{
		Message: message,
		Errors:  errors,
	}

	json.NewEncoder(w).Encode(response)
}

func HandleValidationError(w http.ResponseWriter, validationErrors []string) {
	HandleErrorResponse(w, http.StatusBadRequest, "Validation failed", validationErrors)
}
