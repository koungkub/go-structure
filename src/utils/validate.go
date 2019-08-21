package utils

import (
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (v *CustomValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
