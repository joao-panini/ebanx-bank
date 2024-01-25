package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	erro "github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (accountUseCase *accountUseCase) Transfer(accountOriginID, accountDestinationID, amount int) (*entities.Account, *entities.Account, error) {
	if amount <= 0 {
		return &entities.Account{}, &entities.Account{}, erro.ErrInvalidAmount
	}

	accountOrigin, err := accountUseCase.accountStore.Get(accountOriginID)
	if err != nil {
		return &entities.Account{}, &entities.Account{}, erro.ErrOriginAccNotFound
	}

	if accountOrigin.Balance < amount {
		return &entities.Account{}, &entities.Account{}, erro.ErrInsufficientFunds
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
