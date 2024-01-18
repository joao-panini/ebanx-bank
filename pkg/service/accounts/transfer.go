package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (s *accountService) Transfer(originAccID, destinationAccID, amount int) (*entities.Account, *entities.Account, error) {
	if amount <= 0 {
		return &entities.Account{}, &entities.Account{}, errors.ErrInvalidAmount
	}

	originAcc, err := s.accStore.Get(originAccID)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, errors.ErrOriginAccNotFound
	}

	destAcc, err := s.accStore.Get(destinationAccID)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, errors.ErrDestAccNotFound
	}

	if originAcc.Balance < amount {
		return &entities.Account{}, &entities.Account{}, errors.ErrInsufficientFunds
	}

	originAcc.Balance -= amount
	destAcc.Balance += amount

	s.accStore.Save(originAcc)
	s.accStore.Save(destAcc)
	return originAcc, destAcc, nil
}
