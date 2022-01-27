package utils

import (
	"strings"

	"github.com/google/uuid"
)

func Get_uuid_64() string {
	uuid := uuid.New()
	return uuid.String()
}

func Get_uuid_32() string {
	uuid := Get_uuid_64()
	return strings.Replace(uuid, "-", "", -1)
}
