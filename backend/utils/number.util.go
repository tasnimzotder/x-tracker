package utils

import (
	"github.com/google/uuid"
	"strconv"
)

func ParseInt(id string) (int, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func GenerateUUID() uuid.UUID {
	id := uuid.New()
	return id
}
