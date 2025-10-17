package auth

import (
	"errors"

	"crud/internal/user"
	"crud/pkg/utils"
)

func AuthenticateUserService(email, password string) (string, error) {
	user, err := user.FindUserByEmail(email)
	if err != nil {
		return "", errors.New(err.Error())
	}

	if user == nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(email)
	return token, err
}
