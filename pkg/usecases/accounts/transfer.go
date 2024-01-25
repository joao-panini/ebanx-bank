package accounts

import (
	"github.com/joao-panini/ebanx-bank/pkg/entities"
	erro "github.com/joao-panini/ebanx-bank/pkg/errors"
)

func (accountUseCase *accountUseCase) Transfer(accountOriginID, accountDestinationID, amount int) (*entities.Account, *entities.Account, error) {
	if amount <= 0 {
		return &entities.Account{}, &entities.Account{}, erro.ErrInvalidAmount
	}

	accountOrigin, err := accountUseCase.accountStore.Get(accountOriginID)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, erro.ErrOriginAccNotFound
	}

	accountOrigin, err = accountUseCase.Withdraw(accountOrigin.ID, amount)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, err
	}

	accountDestination, err := accountUseCase.Deposit(accountDestinationID, amount)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, err
	}

	return accountOrigin, accountDestination, nil

}
