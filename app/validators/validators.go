package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}

func ValidateLong(field validator.FieldLevel) bool {
	return len(field.Field().String()) > 12
}

func ValidateLegalTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "legal")
}
