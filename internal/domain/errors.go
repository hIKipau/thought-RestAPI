package domain

import "errors"

var (
	ErrThoughtNotFound = errors.New("thought not found")
	ErrInvalidInput    = errors.New("invalid input")
)
