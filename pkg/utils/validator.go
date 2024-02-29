package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ParseValidatorErr map validator error to list of string
func ParseValidatorErr(err error) []string {
	errMessages := []string{}

	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		errMessages = append(errMessages, fmt.Sprintf("%s failed on %s", e.Field(), e.Tag()))
	}

	return errMessages
}
