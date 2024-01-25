package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	erros "github.com/joao-panini/ebanx-bank/pkg/errors"
)

// Handler para a rota /balance
func (accountHandler *accountHandler) BalanceHandler(writer http.ResponseWriter, request *http.Request) {
	// Get account_id from URL
	accountIDStr := request.URL.Query().Get("account_id")
	if accountIDStr == "" {
		response := erros.ErrAccountIDParamRequired
		http.Error(writer, response.Error(), BadRequest)
		return
	}

	//Convert account_id to int
	accountIDInt, err := strconv.Atoi(accountIDStr)
	if err != nil {
		response := erros.ErrAccountIDParamInvalid
		http.Error(writer, response.Error(), BadRequest)
		return
	}

	//Response needs to be in plain text
	writer.Header().Set(ContentType, PlainTextContentType)
	var response int

	//Call account use case
	account, err := accountHandler.accountUseCase.GetBalance(accountIDInt)
	if err != nil {
		log.Printf("failed get balance: %s\n", err.Error())
		switch {
		case errors.Is(err, erros.ErrOriginAccNotFound):
			http.Error(writer, strconv.Itoa(defaultErrorResponse), NotFound)
		default:
			http.Error(writer, strconv.Itoa(defaultErrorResponse), InternalServerError)
		}
		return
	}

	response = account.Balance
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, response)
}
