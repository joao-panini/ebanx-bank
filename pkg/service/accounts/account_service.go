package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/store"
)

type AccountService interface {
	CreateOrUpdateAccount(id int, initialBalance int) (entities.Account, error)
}

type accountService struct {
	accStore store.AccountStore
}

func NewAccountService(store store.AccountStore) *accountService {
	return &accountService{
		accStore: store,
	}
}
