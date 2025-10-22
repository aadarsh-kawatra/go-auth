package utils

import (
	"regexp"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var EmailRegex = `^[\w.%+-]+@[a-zA-Z0-9\.\-]+\.[a-zA-Z]{2,}$`

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
