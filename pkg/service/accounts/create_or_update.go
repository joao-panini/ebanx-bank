package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (s *accountService) CreateOrUpdateAccount(id int, balance int) (*entities.Account, error) {

	if balance < 0 {
		return &entities.Account{}, errors.ErrInvalidAmount
	}

	//check if id already exists
	acc, err := s.accStore.Get(id)
	if err != nil {
		//Account not found, create new one
		newAcc := &entities.Account{
			ID:      id,
			Balance: balance,
		}
		savedAcc, err := s.accStore.Save(newAcc)
		if err != nil {
			return &entities.Account{}, err
		}
		return savedAcc, nil
	}

	updatedAcc, err := s.Deposit(acc.ID, balance)
	if err != nil {
		return &entities.Account{}, err
	}

	return updatedAcc, nil
}
