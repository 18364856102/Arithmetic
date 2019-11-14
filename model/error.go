package model

import "errors"

var (
	ErrorBadRequest = errors.New("invalid request parameter")
	ErrInvalidRequestType = errors.New("invalid request type")
)
