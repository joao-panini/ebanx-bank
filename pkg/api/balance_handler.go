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

func (h *handler) BalanceHandler(w http.ResponseWriter, r *http.Request) {

	accountIDStr := r.URL.Query().Get("account_id")
	if accountIDStr == "" {
		response := errors.New("missing account_id url parameter")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.Error())
		return
	}
	accountIDInt, err := strconv.Atoi(accountIDStr)
	if err != nil {
		response := errors.New("invalid account_id parameter")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.Error())
		return
	}
	w.Header().Set(ContentType, PlainTextContentType)
	account, err := h.accService.GetBalance(accountIDInt)
	if err != nil {
		log.Printf("failed get balance: %s\n", err.Error())
		switch {
		case errors.Is(err, erro.ErrOriginAccNotFound):
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, 0)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, 0)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, account.Balance)
}
