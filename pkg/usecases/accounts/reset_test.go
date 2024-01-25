package accounts_test

import (
	"errors"
	"testing"

	"github.com/joao-panini/ebanx-bank/pkg/entities"
	"github.com/joao-panini/ebanx-bank/pkg/store"
	"github.com/joao-panini/ebanx-bank/pkg/usecases/accounts"
	"github.com/stretchr/testify/assert"
)

func TestResetUseCase(t *testing.T) {
	genericErr := errors.New("any error")

	type fields struct {
		store store.AccountStore
	}

	testCases := []struct {
		name      string
		fields    fields
		runBefore func(store store.AccountStore)
		wantErr   error
	}{
		{
			name: "usecase should return nil err and reset with store",
			fields: fields{
				store: store.NewAccountStore(),
			},
			runBefore: func(store store.AccountStore) {
				accountTest1 := &entities.Account{
					ID:      1,
					Balance: 100,
				}
				accountTest2 := &entities.Account{
					ID:      2,
					Balance: 200,
				}
				store.Save(accountTest1)
				store.Save(accountTest2)
			},
			wantErr: nil,
		},
		{
			name: "usecase should return error when generic err",
			fields: fields{
				store: store.NewAccountStore(),
			},
			runBefore: func(store store.AccountStore) {
				accountTest1 := &entities.Account{
					ID:      1,
					Balance: 100,
				}
				accountTest2 := &entities.Account{
					ID:      2,
					Balance: 200,
				}
				store.Save(accountTest1)
				store.Save(accountTest2)
			},
			//implement mocks in the future
			wantErr: genericErr,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.runBefore != nil {
				tt.runBefore(tt.fields.store)
			}
			useCase := accounts.NewAccountUseCase(tt.fields.store)

			err := useCase.ResetAccountStates()
			if tt.wantErr != nil {
				if genericErr == tt.wantErr {
					err = genericErr
				}
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			}
		},
		)
	}

}
