package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (s *accountService) Withdraw(accountID int, amount int) (*entities.Account, error) {

	account, err := s.accStore.Get(accountID)
	if err != nil {
		return &entities.Account{}, errors.ErrOriginAccNotFound
	}

	if account.Balance < amount {
		return &entities.Account{}, errors.ErrInsufficientFunds
	}

	account.Balance -= amount
	account = &entities.Account{
		ID:      account.ID,
		Balance: account.Balance,
	}
	updatedAcc, err := s.accStore.Save(account)
	if err != nil {
		return &entities.Account{}, err
	}
	return updatedAcc, nil
}
