package errors

import "errors"

var (
	ErrAnimalNotFound      = errors.New("animal not found")
	ErrInvalidAnimalId     = errors.New("invalid animal id")
	ErrAnimalAlreadyExists = errors.New("animal already exists")
)
