package store

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/errors"
)

func (accountStore *accountStore) Get(id int) (*entities.Account, error) {
	accountStore.mu.Lock()
	defer accountStore.mu.Unlock()
	for _, account := range accountStore.accountStore {
		if account.ID == id {
			return account, nil
		}
	}
	return &entities.Account{}, errors.ErrIdNotFound
}
