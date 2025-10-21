package auth

import (
	"errors"

	"crud/internal/user"
	"crud/pkg/utils"
)

func RegisterUserService(firstName, lastName, email, password string) (string, error) {
	newUser := user.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	err := user.CheckUserExists(newUser)
	if err != nil {
		return "", errors.New(err.Error())
	}

	userId, err := user.CreateUser(newUser)
	if err != nil {
		return "", errors.New(err.Error())
	}

	token, err := utils.GenerateToken(userId, email)
	return token, err
}

func AuthenticateUserService(email, password string) (string, error) {
	user, err := user.FindUserByEmail(email)
	if err != nil {
		return "", errors.New(err.Error())
	}

	if user == nil || !utils.VerifyHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID.Hex(), email)
	return token, err
}
