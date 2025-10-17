package services

import (
	"errors"

	"crud/utils"
)

var dummyUser = map[string]string{
	"admin@example.com": "Admin@123",
}

func AuthenticateUserService(email, password string) (string, error) {
	pass, ok := dummyUser[email]
	if !ok || pass != password {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(email)
	return token, err
}
