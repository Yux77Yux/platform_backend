package utils

import (
	"regexp"
)

func CheckString(str, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}
