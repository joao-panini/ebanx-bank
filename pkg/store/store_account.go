package store

import (
	"errors"
	"sync"

	"github.com/joao-panini/banking-ebanx/pkg/entities"
	erro "github.com/joao-panini/banking-ebanx/pkg/errors"
)

var (
	ErrEmptyID    = errors.New("ID cannot be empty")
	ErrIdNotFound = errors.New("ID not found")
)

type AccountStore interface {
	Save(*entities.Account) (*entities.Account, error)
	Get(id int) (*entities.Account, error)
}

type accountStore struct {
	mu           sync.RWMutex
	accountStore map[int]*entities.Account
}

func NewAccountStore() *accountStore {
	as := make(map[int]*entities.Account)
	return &accountStore{
		accountStore: as,
	}
}

func (s *accountStore) Save(account *entities.Account) (*entities.Account, error) {
	// Should implement database transaction here, to ensure that if any error occurs, all changes are rolled back
	s.mu.Lock()
	defer s.mu.Unlock()
	//verificando por erros neste layer também para simular erros do banco de dados
	if account.ID == 0 {
		return &entities.Account{}, ErrEmptyID
	}

	s.accountStore[account.ID] = account
	return account, nil
}

func (s *accountStore) Get(id int) (*entities.Account, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, a := range s.accountStore {
		if a.ID == id {
			return a, nil
		}
	}
	return &entities.Account{}, erro.ErrIdNotFound
}
