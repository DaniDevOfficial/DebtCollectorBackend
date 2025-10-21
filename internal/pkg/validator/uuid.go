package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func IsValidUUID(fl validator.FieldLevel) bool {
	_, err := uuid.Parse(fl.Field().String())
	return err == nil
}
