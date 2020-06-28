package validator

import (
	"time"

	"github.com/google/uuid"
	validatorV9 "gopkg.in/go-playground/validator.v9"
)

var validator *validatorV9.Validate

func init() {
	validator = validatorV9.New()
	_ = validator.RegisterValidation("uuid", isUUID)
	_ = validator.RegisterValidation("datetime", isDatetime)
}

func Process(i interface{}) error {
	return validator.Struct(i)
}

func isUUID(f validatorV9.FieldLevel) bool {
	v := f.Field().String()
	_, err := uuid.Parse(v)
	return err == nil
}

func isDatetime(f validatorV9.FieldLevel) bool {
	v := f.Field().String()
	_, err := time.Parse(time.RFC3339, v)
	return err == nil
}
