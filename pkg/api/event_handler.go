package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	erro "github.com/joao-panini/banking-ebanx/pkg/errors"
)

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
	case "transfer":
		handleTransfer(h, w, req)
	case "withdraw":
		handleWithdraw(h, w, req)

	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(0)
	}
}

func handleDeposit(h *handler, w http.ResponseWriter, req EventRequest) {

	account, err := h.accService.CreateOrUpdateAccount(req.DestAccountID, int(req.Amount))
	if err != nil {
		//never going to return err
		return
	}

	var res EventResponse
	res.DestinationAcc = AccountResponse{
		ID:      account.ID,
		Balance: account.Balance,
	}

	w.Header().Set(ContentType, JSONContentType)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func handleTransfer(h *handler, w http.ResponseWriter, req EventRequest) {

	origin, dest, err := h.accService.Transfer(req.AccountOriginID, req.DestAccountID, int(req.Amount))
	if err != nil {
		log.Printf("failed make transfer: %s\n", err.Error())
		switch {
		case errors.Is(err, erro.ErrInvalidAmount):
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(0)
		case errors.Is(err, erro.ErrDestAccNotFound):
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(0)
		case errors.Is(err, erro.ErrInsufficientFunds):
			w.WriteHeader(http.StatusPaymentRequired)
			json.NewEncoder(w).Encode(0)
		case errors.Is(err, erro.ErrOriginAccNotFound):
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(0)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(0)
		}
		return
	}
	var res EventResponse
	res.DestinationAcc = AccountResponse{
		ID:      dest.ID,
		Balance: dest.Balance,
	}
	res.OriginAcc = AccountResponse{
		ID:      origin.ID,
		Balance: origin.Balance,
	}
	w.Header().Set(ContentType, JSONContentType)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func handleWithdraw(h *handler, w http.ResponseWriter, req EventRequest) {

	originAcc, err := h.accService.Withdraw(req.AccountOriginID, int(req.Amount))
	if err != nil {
		log.Printf("failed make withdraw: %s\n", err.Error())
		switch {
		case errors.Is(err, erro.ErrInvalidAmount):
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(0)
		case errors.Is(err, erro.ErrInsufficientFunds):
			w.WriteHeader(http.StatusPaymentRequired)
			json.NewEncoder(w).Encode(0)
		case errors.Is(err, erro.ErrOriginAccNotFound):
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(0)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(0)
		}
		return
	}
	var res EventResponse

	res.OriginAcc = AccountResponse{
		ID:      originAcc.ID,
		Balance: originAcc.Balance,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
