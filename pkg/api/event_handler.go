package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	erro "github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (h *handler) EventHandler(w http.ResponseWriter, r *http.Request) {
	var req EventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	err = validateEventRequestByType(EventType(req.Type), h, w, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
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
	accountDestIDInt, err := strconv.Atoi(req.AccountDestId)
	if err != nil {
		return
	}

	account, err := h.accService.CreateOrUpdateAccount(accountDestIDInt, int(req.Amount))
	if err != nil {
		//never going to return err
		return
	}

	var res EventResponse
	res.DestinationAcc = &AccountResponse{
		ID:      account.ID,
		Balance: account.Balance,
	}

	w.Header().Set(ContentType, JSONContentType)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func handleTransfer(h *handler, w http.ResponseWriter, req EventRequest) {
	accOriginIDInt, err := strconv.Atoi(req.AccountOriginID)
	if err != nil {
		return
	}
	accDestIDInt, err := strconv.Atoi(req.AccountDestId)
	if err != nil {
		return
	}
	origin, dest, err := h.accService.Transfer(accOriginIDInt, accDestIDInt, int(req.Amount))
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
	res.DestinationAcc = &AccountResponse{
		ID:      dest.ID,
		Balance: dest.Balance,
	}
	res.OriginAcc = &AccountResponse{
		ID:      origin.ID,
		Balance: origin.Balance,
	}
	w.Header().Set(ContentType, JSONContentType)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func handleWithdraw(h *handler, w http.ResponseWriter, req EventRequest) {
	accOriginIDInt, err := strconv.Atoi(req.AccountOriginID)
	if err != nil {
		return
	}
	originAcc, err := h.accService.Withdraw(accOriginIDInt, int(req.Amount))
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

	res.OriginAcc = &AccountResponse{
		ID:      originAcc.ID,
		Balance: originAcc.Balance,
	}
	w.Header().Set(ContentType, JSONContentType)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func validateEventRequestByType(tipo EventType, h *handler, w http.ResponseWriter, req EventRequest) error {
	if EventType(tipo) != Deposit && EventType(tipo) != Withdraw && EventType(tipo) != Transfer {
		response := fmt.Errorf("type %v not valid", tipo)
		return response
	}

	if EventType(tipo) == Deposit {
		if req.AccountDestId == "" {
			response := fmt.Errorf("destination field required for event type :%v", tipo)
			return response
		}
		if req.Amount == 0 {
			response := fmt.Errorf("amount field required for event type :%v", tipo)
			return response
		}
		if req.AccountOriginID != "" {
			response := fmt.Errorf("origin field not required for event type :%v", tipo)
			return response
		}
	}

	if EventType(tipo) == Transfer {
		if req.AccountOriginID == "" {
			response := fmt.Errorf("origin field required for event type :%v", tipo)
			return response
		}
		if req.AccountDestId == "" {
			response := fmt.Errorf("destination field required for event type :%v", tipo)
			return response
		}
		if req.Amount == 0 {
			response := fmt.Errorf("amount field required for event type :%v", tipo)
			return response
		}
	}

	if EventType(tipo) == Withdraw {
		if req.AccountOriginID == "" {
			response := fmt.Errorf("origin field required for event type :%v", tipo)
			return response
		}
		if req.Amount == 0 {
			response := fmt.Errorf("amount field required for event type :%v", tipo)
			return response
		}
		if req.AccountDestId != "" {
			response := fmt.Errorf("origin field not required for event type :%v", tipo)
			return response
		}
	}
	return nil
}
