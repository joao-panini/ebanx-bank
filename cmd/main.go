package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joao-panini/banking-ebanx/pkg/api"
	"github.com/joao-panini/banking-ebanx/pkg/service/accounts"
	"github.com/joao-panini/banking-ebanx/pkg/store"
)

func main() {
	router := mux.NewRouter()
	store := store.NewAccountStore()
	service := accounts.NewAccountService(store)
	handlers := api.NewHandler(service)

	handlers.SetupRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
