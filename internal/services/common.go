package services

import "errors"

var (
	ErrEventAlreadyPassed = errors.New("event date has already passed")
)
