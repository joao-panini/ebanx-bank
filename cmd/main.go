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
	//
	accountStore := store.NewAccountStore()
	accountUseCase := accounts.NewAccountUseCase(accountStore)
	accountHandlers := api.NewHandler(accountUseCase)

	accountHandlers.SetupRoutes(accountRouter)
	log.Fatal(http.ListenAndServe(":8080", accountRouter))
}
