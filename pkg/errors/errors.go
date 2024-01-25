package errors

import "errors"

var (
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrOriginAccNotFound = errors.New("account origin id not found")
	ErrDestAccNotFound   = errors.New("account destination id not found")
	ErrIdNotFound        = errors.New("account id not found")

	ErrInvalidEventType       = errors.New("invalid event type")
	ErrAccountIDParamRequired = errors.New("missing account_id url parameter")
	ErrAccountIDParamInvalid  = errors.New("invalid account_id parameter")
	ErrTypeNotValid           = errors.New("invalid operation type")
)
