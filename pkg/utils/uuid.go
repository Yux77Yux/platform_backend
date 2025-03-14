package utils

import "github.com/google/uuid"

func GetUuidString() string {
	return uuid.NewString()
}

func GetUuid() uuid.UUID {
	return uuid.New()
}
