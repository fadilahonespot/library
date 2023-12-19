package utility

import (
	"regexp"
	"strings"
)

func CorrectPhoneNumber(phoneNumber string) string {
	if strings.HasPrefix(phoneNumber, "0") {
		phoneNumber = "62" + phoneNumber[1:]
	}
	return phoneNumber
}

func CheckEmail(email string) bool {
	pattern := `^[\w\.-]+@[\w\.-]+\.\w+$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}
