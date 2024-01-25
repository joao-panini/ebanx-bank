package usecases

import "github.com/joao-panini/ebanx-bank/pkg/entities"

type UseCase interface {
	Deposit(destAccID, amount int) (*entities.Account, error)
	Transfer(originAccID, destAccID, amount int) (*entities.Account, *entities.Account, error)
	Withdraw(destAccID, amount int) (*entities.Account, error)
	GetBalance(id int) (*entities.Account, error)
	ResetAccountStates() error
}
