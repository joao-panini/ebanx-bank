package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (s *accountService) GetBalance(accountID int) (*entities.Account, error) {
	account, err := s.accStore.Get(accountID)
	if err != nil {
		return &entities.Account{}, errors.ErrOriginAccNotFound
	}
	return account, nil
}
