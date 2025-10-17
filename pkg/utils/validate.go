package utils

import "regexp"

var EmailRegex = `^[\w.%+-]+@[a-zA-Z0-9\.\-]+\.[a-zA-Z]{2,}$`

func IsEmail(email string) bool {
	matched, err := regexp.MatchString(EmailRegex, email)
	if err != nil {
		return false
	}
	return matched
}
