package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"crud/infrastructure/middlewares"
	"crud/pkg/utils"
)

func GetProfileHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	userId := parts[2]
	if !utils.IsValidObjectID(userId) {
		http.Error(w, "Invalid User Id", http.StatusBadRequest)
		return
	}

	claims := middlewares.GetUserFromContext(r)

	err := ValidateUserAccess(userId, claims.Id)
	if err != nil {
		var errors []string
		utils.HandleErrorResponse(w, http.StatusUnauthorized, "Error", append(errors, err.Error()))
		return
	}

	user, err := GetUserProfileService(userId)
	if err != nil {
		var errors []string
		utils.HandleErrorResponse(w, http.StatusBadRequest, "Error", append(errors, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(GetProfileResponse{
		Code:    http.StatusOK,
		Message: "Success",
		User:    user,
	})
}
