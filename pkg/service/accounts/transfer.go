package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	erro "github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (s *accountService) Transfer(originAccID, destinationAccID, amount int) (*entities.Account, *entities.Account, error) {
	if amount <= 0 {
		return &entities.Account{}, &entities.Account{}, erro.ErrInvalidAmount
	}

	originAcc, err := s.accStore.Get(originAccID)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, erro.ErrOriginAccNotFound
	}

	if originAcc.Balance < amount {
		return &entities.Account{}, &entities.Account{}, err
	}

	originAcc, err = s.Withdraw(originAcc.ID, amount)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, err
	}

	destAcc, err := s.Deposit(destinationAccID, amount)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, err
	}

	return originAcc, destAcc, nil

}
