package store

import (
	"errors"
	"sync"

	"github.com/joao-panini/banking-ebanx/pkg/entities"
)

var (
	ErrEmptyID    = errors.New("ID cannot be empty")
	ErrIdNotFound = errors.New("ID not found")
)

type AccountStore interface {
	Save(*entities.Account) (*entities.Account, error)
	Get(id int) (*entities.Account, error)
	ResetAccountStates() error
}

type accountStore struct {
	mu           sync.RWMutex
	accountStore map[int]*entities.Account
}

func NewAccountStore() *accountStore {
	acountStore := make(map[int]*entities.Account)
	return &accountStore{
		accountStore: acountStore,
	}
}
