package util

import "github.com/google/uuid"

func GetUUID() string {
	return genUUID()
}

func genUUID() string {
	return uuid.New().String()
}
