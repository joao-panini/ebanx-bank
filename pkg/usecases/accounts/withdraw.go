package accounts

import (
	"github.com/joao-panini/ebanx-bank/pkg/entities"
	"github.com/joao-panini/ebanx-bank/pkg/errors"
)

func (accountUseCase *accountUseCase) Withdraw(accountID int, amount int) (*entities.Account, error) {
	if amount < 0 {
		return &entities.Account{}, errors.ErrInvalidAmount
	}
	account, err := accountUseCase.accountStore.Get(accountID)
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
	updatedAcc, err := accountUseCase.accountStore.Save(account)
	if err != nil {
		return &entities.Account{}, err
	}
	return updatedAcc, nil
}
