package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GetUUID64() string {
	uuid := uuid.New()
	return uuid.String()
}

func GetUUID32() string {
	uuid := GetUUID64()
	return strings.Replace(uuid, "-", "", -1)
}
