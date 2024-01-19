package api

type EventType string

const (
	Deposit  EventType = "deposit"
	Transfer EventType = "transfer"
	Withdraw EventType = "withdraw"
)

var ValidEventTypes = [3]EventType{Deposit, Transfer, Withdraw}

type EventRequest struct {
	Type            string `json:"type" validate:"required,min=1"`
	AccountOriginID string `json:"origin"`
	AccountDestId   string `json:"destination"`
	Amount          int    `json:"amount" validate:"required,min=0"`
}

type EventResponse struct {
	DestinationAcc *AccountResponse `json:"destination,omitempty"`
	OriginAcc      *AccountResponse `json:"origin,omitempty"`
}

type AccountResponse struct {
	ID      int `json:"id,omitempty"`
	Balance int `json:"balance,omitempty"`
}
