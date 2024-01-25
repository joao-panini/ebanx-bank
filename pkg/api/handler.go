package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joao-panini/ebanx-bank/pkg/usecases/accounts"
)

const (
	ContentType          = "Content-Type"
	JSONContentType      = "application/json"
	PlainTextContentType = "text/plain"
	DateLayout           = "2006-01-02T15:04:05Z"
)

type AccountHandler interface {
	ResetHandler(writer http.ResponseWriter, request *http.Request)
	EventHandler(writer http.ResponseWriter, request *http.Request)
	BalanceHandler(writer http.ResponseWriter, request *http.Request)
}

type accountHandler struct {
	accountUseCase accounts.UseCase
}

func NewAccountHandler(accountUseCase accounts.UseCase) *accountHandler {
	return &accountHandler{
		accountUseCase: accountUseCase,
	}
}

func (accountHandler *accountHandler) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/reset", accountHandler.ResetHandler).Methods("POST")
	router.HandleFunc("/event", accountHandler.EventHandler).Methods("POST")
	router.HandleFunc("/balance", accountHandler.BalanceHandler).Methods("GET")
}
