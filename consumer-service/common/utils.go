package common

import (
	errorv2 "consumer-service/errors"
	"strings"

	"github.com/google/uuid"
)

func ValidateUUID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errorv2.BadRequest.New("uuid cannot be empty")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return errorv2.BadRequest.New("invalid uuid format")
	}

	return nil
}
