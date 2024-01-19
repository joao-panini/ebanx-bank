package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joao-panini/banking-ebanx/pkg/service/accounts"
)

const (
	ContentType          = "Content-Type"
	JSONContentType      = "application/json"
	PlainTextContentType = "text/plain"
	DateLayout           = "2006-01-02T15:04:05Z"
)

type Handler interface {
	ResetHandler(w http.ResponseWriter, r *http.Request)
	EventHandler(w http.ResponseWriter, r *http.Request)
	BalanceHandler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	accService accounts.AccountService
}

func NewHandler(accountService accounts.AccountService) *handler {
	return &handler{
		accService: accountService,
	}
}

func (h *handler) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/reset", h.ResetHandler).Methods("POST")
	router.HandleFunc("/event", h.EventHandler).Methods("POST")
	router.HandleFunc("/balance", h.BalanceHandler).Methods("GET")
}
