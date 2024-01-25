package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (accountUseCase *accountUseCase) Deposit(accountID int, amount int) (*entities.Account, error) {
	if amount < 0 {
		return &entities.Account{}, errors.ErrInvalidAmount
	}

	account, err := accountUseCase.accountStore.Get(accountID)
	if err != nil {
		//Account not found, create new one
		newAcc := &entities.Account{
			ID:      accountID,
			Balance: amount,
		}
		savedAcc, err := accountUseCase.accountStore.Save(newAcc)
		if err != nil {
			return &entities.Account{}, err
		}
		return savedAcc, nil
	}

	account.Balance += amount
	updatedAcc, err := accountUseCase.accountStore.Save(account)
	if err != nil {
		return &entities.Account{}, err
	}
	return updatedAcc, nil
}
