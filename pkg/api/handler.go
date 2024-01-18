package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joao-panini/banking-ebanx/pkg/service/accounts"
)

type Handler interface {
	ResetHandler(w http.ResponseWriter, r *http.Request)
	EventHandler(w http.ResponseWriter, r *http.Request)
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
}

func (h *handler) ResetHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}