package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (accountUseCase *accountUseCase) GetBalance(accountID int) (*entities.Account, error) {
	account, err := accountUseCase.accountStore.Get(accountID)
	if err != nil {
		return &entities.Account{}, errors.ErrOriginAccNotFound
	}
	return account, nil
}
