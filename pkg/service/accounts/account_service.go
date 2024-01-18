package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/store"
)

type AccountService interface {
	CreateOrUpdateAccount(id int, amount int) (*entities.Account, error)
	Deposit(destAccID, amount int) (*entities.Account, error)
	Transfer(originAccID, destAccID, amount int) (*entities.Account, *entities.Account, error)
	Withdraw(destAccID, amount int) (*entities.Account, error)
	GetBalance(id int) (*entities.Account, error)
	ResetAccountStates() error
}

type accountService struct {
	accStore store.AccountStore
}

func NewAccountService(store store.AccountStore) *accountService {
	return &accountService{
		accStore: store,
	}
}
