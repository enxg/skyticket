package validator

import "github.com/go-playground/validator/v10"

type StructValidator interface {
	Validate(out any) error
}

type structValidator struct {
	validate *validator.Validate
}

func NewStructValidator() StructValidator {
	return &structValidator{
		validate: validator.New(),
	}
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}
