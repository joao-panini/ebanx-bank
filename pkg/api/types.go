package api

type EventRequest struct {
	Type            string `json:"type" validate:"required,min=1"`
	AccountOriginID int    `json:"origin"`
	DestAccountID   int    `json:"destination"`
	Amount          int    `json:"amount" validate:"required,min=0"`
}

type EventResponse struct {
	DestinationAcc AccountResponse `json:"destination,omitempty"`
	OriginAcc      AccountResponse `json:"origin,omitempty"`
}

type AccountResponse struct {
	ID      int `json:"id,omitempty"`
	Balance int `json:"balance,omitempty"`
}

type BalanceRequest struct {
}