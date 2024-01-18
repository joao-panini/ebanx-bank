package accounts

import (
	"errors"

	"github.com/joao-panini/banking-ebanx/pkg/entities"
)

var (
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

func (s *accountService) CreateOrUpdateAccount(id int, balance int) (entities.Account, error) {

	if balance < 0 {
		return entities.Account{}, ErrInvalidAmount
	}

	//check if id already exists
	acc, err := s.accStore.Get(id)
	if err != nil {
		//Account not found, create new one
		newAcc := entities.Account{
			ID:      id,
			Balance: balance,
		}
		s.accStore.Save(newAcc)
		return newAcc, nil
	}

	s.Deposit(acc.ID, balance)

	return acc, nil
}

func (s *accountService) Deposit(accountID int, amount int) (entities.Account, error) {
	if amount <= 0 {
		return entities.Account{}, ErrInvalidAmount
	}

	account, err := s.accStore.Get(accountID)
	if err != nil {
		return entities.Account{}, err
	}

	account.Balance += amount
	s.accStore.Save(account)
	return account, nil
}
