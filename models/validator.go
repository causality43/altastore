package models

import (
	validator "github.com/go-playground/validator/v10"
)

var Validate *validator.Validate = validator.New()
