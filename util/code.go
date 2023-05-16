package util

import "strings"

func StandarizeEmail(email string) string {
	return strings.ToLower(email)
}
