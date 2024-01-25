package store

import "github.com/joao-panini/ebanx-bank/pkg/entities"

func (accountStore *accountStore) ResetAccountStates() error {
	accountStore.mu.Lock()
	defer accountStore.mu.Unlock()

	accountStore.accountStore = make(map[int]*entities.Account)
	return nil
}
