package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joao-panini/ebanx-bank/pkg/api"
	"github.com/joao-panini/ebanx-bank/pkg/store"
	"github.com/joao-panini/ebanx-bank/pkg/usecases/accounts"
)

func main() {
	// Instancia o router
	accountRouter := mux.NewRouter()
	// Instancia uma store de account
	accountStore := store.NewAccountStore()
	// Instancia UseCases de account
	accountUseCase := accounts.NewAccountUseCase(accountStore)
	// Inicia Handlers de account
	accountHandlers := api.NewAccountHandler(accountUseCase)

	accountHandlers.SetupRoutes(accountRouter)
	log.Fatal(http.ListenAndServe(":8080", accountRouter))
}
