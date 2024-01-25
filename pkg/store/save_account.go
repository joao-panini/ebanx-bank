package store

import "github.com/joao-panini/banking-ebanx/pkg/entities"

func (accountStore *accountStore) Save(account *entities.Account) (*entities.Account, error) {
	// Should implement database transaction here, to ensure that if any error occurs, all changes are rolled back
	accountStore.mu.Lock()
	defer accountStore.mu.Unlock()
	//verificando por erros neste layer também para simular erros do banco de dados
	if account.ID == 0 {
		return &entities.Account{}, ErrEmptyID
	}

	accountStore.accountStore[account.ID] = account
	return account, nil
}
