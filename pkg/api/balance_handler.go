package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	erro "github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (h *handler) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["account_id"]
	accountIDInt, err := strconv.Atoi(accountID)
	if err != nil {
		return
	}
	account, err := h.accService.GetBalance(accountIDInt)
	if err != nil {
		log.Printf("failed make transfer: %s\n", err.Error())
		switch {
		case errors.Is(err, erro.ErrOriginAccNotFound):
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(0)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(0)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account.Balance)
}
