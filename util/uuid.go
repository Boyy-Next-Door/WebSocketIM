package util

import (
	uuid "github.com/satori/go.uuid"
)

func GetUUID() string {
	u2 := uuid.NewV4()
	return u2.String()
}
