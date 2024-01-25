package accounts

import (
	"github.com/joao-panini/ebanx-bank/pkg/entities"
	"github.com/joao-panini/ebanx-bank/pkg/store"
)

type UseCase interface {
	Deposit(destinationAccountID, amount int) (*entities.Account, error)
	Transfer(accountOriginID, destinationAccountID, amount int) (*entities.Account, *entities.Account, error)
	Withdraw(accountOriginID, amount int) (*entities.Account, error)
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
