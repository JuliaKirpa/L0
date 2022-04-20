package pkg

import (
	"NatsMC/Consumer/models"
	"encoding/json"
	"errors"
)

func ValidateMessage(message []byte) error {
	msg := &models.Order{}

	err := json.Unmarshal(message, msg)
	if err != nil {
		return errors.New("incorrect type of message")
	}
	return nil
}
