package utils

import "gopkg.in/go-playground/validator.v9"

var Validate *validator.Validate

// InitValidator initialize validator instance
func InitValidator() {
	Validate = validator.New()
}