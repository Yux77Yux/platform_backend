package utils

import "fmt"

func CheckStringLength(obj string, min, max int) error {
	length := len(obj)
	if length < min || length > max {
		return fmt.Errorf("error: string not match require,require (%d,%d)", min, max)
	}
	return nil
}
