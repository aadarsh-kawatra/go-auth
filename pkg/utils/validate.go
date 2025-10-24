package utils

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var EmailRegex = `^[\w.%+-]+@[a-zA-Z0-9\.\-]+\.[a-zA-Z]{2,}$`

var Validate = validator.New()

func IsEmail(email string) bool {
	matched, err := regexp.MatchString(EmailRegex, email)
	if err != nil {
		return false
	}
	return matched
}

func IsValidObjectID(id string) bool {
	_, err := bson.ObjectIDFromHex(id)
	return err == nil
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", fe.Field(), fe.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}

func ValidateStruct(s any) []string {
	var validationErrors []string

	err := Validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, getErrorMessage(err))
		}
	}

	return validationErrors
}
