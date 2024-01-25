package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	erros "github.com/joao-panini/banking-ebanx/pkg/errors"
)

// Handler da rota /event
func (accountHandler *accountHandler) EventHandler(writer http.ResponseWriter, request *http.Request) {
	var eventRequest EventRequest
	err := json.NewDecoder(request.Body).Decode(&eventRequest)
	if err != nil {
		http.Error(writer, "Failed to parse JSON", BadRequest)
		return
	}
	err = validateEventType(EventType(eventRequest.Type), accountHandler, writer, eventRequest)
	if err != nil {
		http.Error(writer, err.Error(), BadRequest)
		return
	}

	eventResponse, code, err := handleRequestType(accountHandler, writer, eventRequest)
	if err != nil {
		http.Error(writer, strconv.Itoa(defaultErrorResponse), code)
		return
	}

	writer.Header().Set(ContentType, JSONContentType)
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(eventResponse)

}

func handleDeposit(accountHandler *accountHandler, writer http.ResponseWriter, request EventRequest) (EventResponse, int, error) {
	accountDestIDInt, err := strconv.Atoi(request.AccountDestId)
	if err != nil {
		return EventResponse{}, BadRequest, fmt.Errorf("erro convertendo account destination ID de string para int %writer", err)
	}

	account, err := accountHandler.accountUseCase.Deposit(accountDestIDInt, request.Amount)
	if err != nil {
		//never going to return err
		return EventResponse{}, BadRequest, err
	}

	var response EventResponse
	response.DestinationAcc = &AccountResponse{
		ID:      strconv.Itoa(account.ID),
		Balance: account.Balance,
	}

	return response, Created, nil
}

func handleTransfer(accountHandler *accountHandler, writer http.ResponseWriter, req EventRequest) (EventResponse, int, error) {
	accountOriginIDInt, err := strconv.Atoi(req.AccountOriginID)
	if err != nil {
		return EventResponse{}, BadRequest, fmt.Errorf("erro convertendo account origin ID de string para int %writer", err)
	}
	accountDestinationIDInt, err := strconv.Atoi(req.AccountDestId)
	if err != nil {
		return EventResponse{}, BadRequest, fmt.Errorf("erro convertendo account destination ID de string para int %writer", err)
	}

	accountOrigin, accountDestination, err := accountHandler.accountUseCase.Transfer(accountOriginIDInt, accountDestinationIDInt, req.Amount)
	if err != nil {
		log.Printf("failed make transfer: %s\n", err.Error())
		switch {
		case errors.Is(err, erros.ErrInvalidAmount):
			return EventResponse{}, BadRequest, erros.ErrInvalidAmount

		case errors.Is(err, erros.ErrInsufficientFunds):
			return EventResponse{}, http.StatusPaymentRequired, erros.ErrInsufficientFunds

		case errors.Is(err, erros.ErrOriginAccNotFound):
			return EventResponse{}, NotFound, erros.ErrOriginAccNotFound

		default:
			return EventResponse{}, InternalServerError, err
		}
	}

	var response EventResponse
	response.DestinationAcc = &AccountResponse{
		ID:      strconv.Itoa(accountDestination.ID),
		Balance: accountDestination.Balance,
	}
	response.OriginAcc = &AccountResponse{
		ID:      strconv.Itoa(accountOrigin.ID),
		Balance: accountOrigin.Balance,
	}

	return response, Created, nil
}

func handleWithdraw(accountHandler *accountHandler, writer http.ResponseWriter, request EventRequest) (EventResponse, int, error) {
	accountOriginIDInt, err := strconv.Atoi(request.AccountOriginID)
	if err != nil {
		return EventResponse{}, BadRequest, fmt.Errorf("erro convertendo account origin ID de string para int. %writer", err)
	}
	accountOrigin, err := accountHandler.accountUseCase.Withdraw(accountOriginIDInt, int(request.Amount))
	if err != nil {
		log.Printf("failed make withdraw: %s\n", err.Error())
		switch {
		case errors.Is(err, erros.ErrInvalidAmount):
			return EventResponse{}, BadRequest, erros.ErrInvalidAmount

		case errors.Is(err, erros.ErrInsufficientFunds):
			return EventResponse{}, BadRequest, erros.ErrInsufficientFunds

		case errors.Is(err, erros.ErrOriginAccNotFound):
			return EventResponse{}, NotFound, erros.ErrInsufficientFunds
		default:
			return EventResponse{}, InternalServerError, err
		}
	}
	var response EventResponse

	response.OriginAcc = &AccountResponse{
		ID:      strconv.Itoa(accountOrigin.ID),
		Balance: accountOrigin.Balance,
	}

	return response, Created, nil
}

func handleRequestType(accountHandler *accountHandler, writer http.ResponseWriter, request EventRequest) (EventResponse, int, error) {
	switch request.Type {
	case "deposit":
		response, code, err := handleDeposit(accountHandler, writer, request)
		if err != nil {
			return EventResponse{}, code, err
		}
		return response, code, nil

	case "transfer":
		response, code, err := handleTransfer(accountHandler, writer, request)
		if err != nil {
			return EventResponse{}, code, err
		}
		return response, code, nil

	case "withdraw":
		response, code, err := handleWithdraw(accountHandler, writer, request)
		if err != nil {
			return EventResponse{}, code, err
		}
		return response, code, nil
	default:
		http.Error(writer, erros.ErrTypeNotValid.Error(), BadRequest)
		return EventResponse{}, defaultErrorResponse, erros.ErrTypeNotValid
	}

}

func validateEventType(tipo EventType, accountHandler *accountHandler, writer http.ResponseWriter, request EventRequest) error {
	if EventType(tipo) != Deposit && EventType(tipo) != Withdraw && EventType(tipo) != Transfer {
		response := fmt.Errorf("type %v not valid", tipo)
		return response
	}

	if EventType(tipo) == Deposit {
		if request.AccountDestId == "" {
			response := fmt.Errorf("destination field required for event type :%v", tipo)
			return response
		}
		if request.Amount == 0 {
			response := fmt.Errorf("amount field required for event type :%v", tipo)
			return response
		}
		if request.AccountOriginID != "" {
			response := fmt.Errorf("origin field not required for event type :%v", tipo)
			return response
		}
	}

	if EventType(tipo) == Transfer {
		if request.AccountOriginID == "" {
			response := fmt.Errorf("origin field required for event type :%v", tipo)
			return response
		}
		if request.AccountDestId == "" {
			response := fmt.Errorf("destination field required for event type :%v", tipo)
			return response
		}
		if request.Amount == 0 {
			response := fmt.Errorf("amount field required for event type :%v", tipo)
			return response
		}
	}

	if EventType(tipo) == Withdraw {
		if request.AccountOriginID == "" {
			response := fmt.Errorf("origin field required for event type :%v", tipo)
			return response
		}
		if request.Amount == 0 {
			response := fmt.Errorf("amount field required for event type :%v", tipo)
			return response
		}
		if request.AccountDestId != "" {
			response := fmt.Errorf("origin field not required for event type :%v", tipo)
			return response
		}
	}
	return nil
}
