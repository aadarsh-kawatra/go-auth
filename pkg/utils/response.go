package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseStruct struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func HandleValidationError(w http.ResponseWriter, validationErrors []string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response := ErrorResponseStruct{
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		Errors:  validationErrors,
	}

	json.NewEncoder(w).Encode(response)
}
