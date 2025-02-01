package utils

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		validatorInstance = validator.New(validator.WithRequiredStructEnabled())
	})
	return validatorInstance
}

func GetFirstValiationErr(err error) string {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return "Invalid validation error"
	}

	for _, err := range err.(validator.ValidationErrors) {
		return fmt.Sprintf("%s must be %s", err.Field(), err.Tag())
	}

	return "Unknown validation error"
}

func DoValidate(data interface{}) (string, error) {
	v := GetValidator()
	err := v.Struct(data)

	if err == nil {
		return "", nil
	}

	return GetFirstValiationErr(err), err
}
