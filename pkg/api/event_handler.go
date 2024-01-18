package api

import (
	"encoding/json"
	"net/http"
)

type EventRequest struct {
	Type      string `json:"event_type" validate:"required"`
	AccountID int    `json:"account_id" validate:"required"`
	Amount    int    `json:"amount" validate:"required"`
}

func (h *handler) EventHandler(w http.ResponseWriter, r *http.Request) {
	var req EventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	switch req.Type {
	case "deposit":
		handleDeposit(h, w, req)
	default:
		http.Error(w, "Invalid event type", http.StatusBadRequest)
	}
}

func handleDeposit(h *handler, w http.ResponseWriter, req EventRequest) {
	if req.AccountID == 0 {
		http.Error(w, "Invalid destination account", http.StatusBadRequest)
	}

	account, err := h.accService.CreateOrUpdateAccount(req.AccountID, int(req.Amount))
	if err != nil {
		http.Error(w, "Error creating or updating account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&account)
}
