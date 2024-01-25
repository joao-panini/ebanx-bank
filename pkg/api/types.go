package api

import (
	"net/http"
)

type EventType string

const (
	Deposit  EventType = "deposit"
	Transfer EventType = "transfer"
	Withdraw EventType = "withdraw"
)

var (
	//HTTP STATUS CODE 400
	BadRequest = http.StatusBadRequest
	//HTTP STATUS CODE 404
	NotFound = http.StatusNotFound
	//HTTP STATUS CODE 500
	InternalServerError = http.StatusInternalServerError
	Created             = http.StatusCreated

	//Test suit requires 0 as int when returning error from account balance
	defaultErrorResponse = 0
	//Tipos validos de evento
	ValidEventTypes = [3]EventType{Deposit, Transfer, Withdraw}
)

type EventRequest struct {
	Type            string `json:"type"`
	AccountOriginID string `json:"origin"`
	AccountDestId   string `json:"destination"`
	Amount          int    `json:"amount" validate:"required,min=0"`
}

type EventResponse struct {
	DestinationAcc *AccountResponse `json:"destination,omitempty"`
	OriginAcc      *AccountResponse `json:"origin,omitempty"`
}

type AccountResponse struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}
