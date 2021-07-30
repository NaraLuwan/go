package main

import (
	"regexp"
)

func CheckEmail(email string) bool {
	const pattern = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
