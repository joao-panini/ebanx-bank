package accounts_test

import (
	"errors"
	"testing"

	"github.com/joao-panini/ebanx-bank/pkg/entities"
	erros "github.com/joao-panini/ebanx-bank/pkg/errors"
	"github.com/joao-panini/ebanx-bank/pkg/store"
	"github.com/joao-panini/ebanx-bank/pkg/usecases/accounts"
	"github.com/stretchr/testify/assert"
)

func TestDepositUseCase(t *testing.T) {
	genericErr := errors.New("any error")

	type fields struct {
		store store.AccountStore
	}
	type args struct {
		accountID int
		amount    int
	}
	testCases := []struct {
		name      string
		fields    fields
		args      args
		runBefore func(args args, store store.AccountStore)
		want      *entities.Account
		wantErr   error
	}{
		{
			name: "should return account with updated balance when trying to deposit to an existing account",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountID: 1,
				amount:    100,
			},
			runBefore: func(args args, store store.AccountStore) {
				accountTest := &entities.Account{
					ID:      1,
					Balance: 100,
				}
				store.Save(accountTest)
			},
			want: &entities.Account{
				ID:      1,
				Balance: 200,
			},
		},
		{
			name: "usecase should create account and return it with the given balance when trying to deposit to a account that does not exists",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountID: 2,
				amount:    50,
			},
			runBefore: func(args args, store store.AccountStore) {
				accountTest := &entities.Account{
					ID:      1,
					Balance: 100,
				}
				store.Save(accountTest)
			},
			want: &entities.Account{
				ID:      2,
				Balance: 50,
			},
		},
		{
			name: "usecase should return ErrInvalidAmount when deposit amount is lower than 0",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountID: 1,
				amount:    -50,
			},
			want:    &entities.Account{},
			wantErr: erros.ErrInvalidAmount,
		},
		{
			name:   "should return error when deposit usecase returns any errors",
			fields: fields{store.NewAccountStore()},
			args:   args{},
			runBefore: func(args args, store store.AccountStore) {
			},
			want:    &entities.Account{},
			wantErr: genericErr,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.store.ResetAccountStates()
			if tt.runBefore != nil {
				tt.runBefore(tt.args, tt.fields.store)
			}
			useCase := accounts.NewAccountUseCase(tt.fields.store)

			got, err := useCase.Deposit(tt.args.accountID, tt.args.amount)
			if tt.wantErr != nil {
				if tt.wantErr == genericErr {
					err = genericErr
				}
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				tt.want.ID = got.ID
				tt.want.Balance = got.Balance
				assert.Equal(t, tt.want, got)
			}
		},
		)
	}

}
