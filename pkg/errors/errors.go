package errors

import "errors"

var (
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrOriginAccNotFound = errors.New("account origin not found")
	ErrDestAccNotFound   = errors.New("destination origin not found")
)
