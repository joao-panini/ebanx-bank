package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (s *accountService) Deposit(accountID int, amount int) (*entities.Account, error) {
	if amount < 0 {
		return &entities.Account{}, errors.ErrInvalidAmount
	}

	account, err := s.accStore.Get(accountID)
	if err != nil {
		//Account not found, create new one
		newAcc := &entities.Account{
			ID:      accountID,
			Balance: amount,
		}
		savedAcc, err := s.accStore.Save(newAcc)
		if err != nil {
			return &entities.Account{}, err
		}
		return savedAcc, nil
	}

	account.Balance += amount
	updatedAcc, err := s.accStore.Save(account)
	if err != nil {
		return &entities.Account{}, err
	}
	return updatedAcc, nil
}
