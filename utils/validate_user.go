package utils

import "regexp"

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	return hasUpper && hasNumber
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
