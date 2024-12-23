package errors

import "errors"

var (
	ErrInvalidJSON       = errors.New("invalid JSON format")
	ErrOddNumberOfPeople = errors.New("the number of participants must be even")
	ErrInternalError     = errors.New("an internal error occurred")
)
