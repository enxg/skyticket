package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type StructValidator interface {
	Validate(out any) error
}

type structValidator struct {
	validate *validator.Validate
}

type ValidationError struct {
	Field string `json:"field" example:"field_name"`
	Error string `json:"error" example:"Error message for the field"`
}

type ValidationErrors = validator.ValidationErrors

func NewStructValidator() StructValidator {
	vld := validator.New()

	vld.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &structValidator{
		validate: vld,
	}
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func ParseValidationErrors(validationErrors validator.ValidationErrors) []ValidationError {
	errs := make([]ValidationError, len(validationErrors))
	for i, ve := range validationErrors {
		errs[i] = ValidationError{
			Field: ve.Field(),
			Error: validationErrorToText(ve),
		}
	}

	return errs
}

func validationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", e.Field())
	case "datetime":
		return fmt.Sprintf("%s must be in RFC3339 format.", e.Field())
	default:
		return fmt.Sprintf("%s is invalid.", e.Field())
	}
}
