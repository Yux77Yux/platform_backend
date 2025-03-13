package utils

import "fmt"

func CheckStringLength(obj string, min, max int) error {
	length := len(obj)
	if length < min || length > max {
		return fmt.Errorf("error: %s not match require,require (%d,%d)", obj, min, max)
	}
	return nil
}
