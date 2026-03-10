package plog

import "errors"

var (
	ErrUnknownEvent      = errors.New("Unknown Event")
	ErrUnprocessableBody = errors.New("Unprocessable body")
)
