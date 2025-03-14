package utils

import (
	"regexp"
)

func CheckString(str, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}

func GetMatches(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindStringSubmatch(str)
}
