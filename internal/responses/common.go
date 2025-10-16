package responses

import "github.com/enxg/skyticket/pkg/validator"

type ValidationErrorResponse struct {
	Errors []validator.ValidationError `json:"errors"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Error message"`
}
