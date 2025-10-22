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

	if claims.Id != userId {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := FindUserById(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user == nil {
		http.Error(w, "User Not Found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(GetProfileResponse{
		Code:    http.StatusOK,
		Message: "Success",
		User:    user,
	})
}
