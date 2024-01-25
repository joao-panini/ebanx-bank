package accounts

import (
	"github.com/joao-panini/banking-ebanx/pkg/entities"
	"github.com/joao-panini/banking-ebanx/pkg/store"
)

type UseCase interface {
	Deposit(destAccID, amount int) (*entities.Account, error)
	Transfer(originAccID, destAccID, amount int) (*entities.Account, *entities.Account, error)
	Withdraw(destAccID, amount int) (*entities.Account, error)
	GetBalance(id int) (*entities.Account, error)
	ResetAccountStates() error
}

type accountUseCase struct {
	accountStore store.AccountStore
}

func NewAccountUseCase(accountStore store.AccountStore) *accountUseCase {
	return &accountUseCase{
		accountStore: accountStore,
	}
}
