package api

import (
	"net/http"
)

type Response struct {
	OK bool
}

// Handler da rota /reset
func (accountHandler *accountHandler) ResetHandler(writer http.ResponseWriter, request *http.Request) {

	err := accountHandler.accountUseCase.ResetAccountStates()
	if err != nil {
		http.Error(writer, err.Error(), InternalServerError)
	}

	writer.Header().Set(ContentType, PlainTextContentType)
	writer.Write([]byte("OK"))
}
