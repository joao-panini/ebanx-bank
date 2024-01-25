package accounts_test

import (
	"testing"

	"github.com/joao-panini/ebanx-bank/pkg/entities"
	"github.com/joao-panini/ebanx-bank/pkg/errors"
	"github.com/joao-panini/ebanx-bank/pkg/store"
	"github.com/joao-panini/ebanx-bank/pkg/usecases/accounts"
	"github.com/stretchr/testify/assert"
)

func TestGetBalanceUseCase(t *testing.T) {

	type fields struct {
		store store.AccountStore
	}
	type args struct {
		accountID int
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
			name: "should return success when account exists when getting balance",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountID: 1,
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
				Balance: 100,
			},
		},
		{
			name: "should return error when account origin does not exist when getting balance",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountID: 1,
			},
			runBefore: func(args args, store store.AccountStore) {
				accountTest := &entities.Account{
					ID:      2,
					Balance: 200,
				}
				store.Save(accountTest)
			},
			want:    &entities.Account{},
			wantErr: errors.ErrOriginAccNotFound,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.store.ResetAccountStates()
			if tt.runBefore != nil {
				tt.runBefore(tt.args, tt.fields.store)
			}
			useCase := accounts.NewAccountUseCase(tt.fields.store)

			got, err := useCase.GetBalance(tt.args.accountID)
			if tt.wantErr != nil {
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
