package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	// Initialize the validator instance
	validate = validator.New()
}

// ValidateStruct validates a struct using the registered validation tags
func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		// Validation failed
		return err
	}

	return nil
}
