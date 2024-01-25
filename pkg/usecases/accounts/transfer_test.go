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

func TestTransferUseCase(t *testing.T) {
	genericErr := errors.New("any error")

	type fields struct {
		store store.AccountStore
	}
	type args struct {
		accountOriginID      int
		accountDestinationID int
		amount               int
	}
	testCases := []struct {
		name      string
		fields    fields
		args      args
		runBefore func(args args, store store.AccountStore)
		want      []*entities.Account
		wantErr   error
	}{
		{
			name: "should return origin and destination account with updated values when transfering from existing account to any destination account",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountOriginID:      1,
				accountDestinationID: 2,
				amount:               50,
			},
			runBefore: func(args args, store store.AccountStore) {
				accountOriginTest := &entities.Account{
					ID:      1,
					Balance: 200,
				}
				accountDestinationTest := &entities.Account{
					ID:      2,
					Balance: 300,
				}
				store.Save(accountOriginTest)
				store.Save(accountDestinationTest)
			},
			want: []*entities.Account{
				{
					ID:      1,
					Balance: 150,
				},
				{
					ID:      2,
					Balance: 150,
				},
			},
			wantErr: nil,
		},
		{
			name: "should return account origin and create new account if destination account doest not exists",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountOriginID:      1,
				accountDestinationID: 2,
				amount:               50,
			},
			runBefore: func(args args, store store.AccountStore) {
				accountOriginTest := &entities.Account{
					ID:      1,
					Balance: 200,
				}

				store.Save(accountOriginTest)
			},
			want: []*entities.Account{
				{
					ID:      1,
					Balance: 150,
				},
				{
					ID:      2,
					Balance: 50,
				},
			},
			wantErr: nil,
		},
		{
			name: "should return ErrOriginAccNotFound when account origin does not exists",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountOriginID:      1,
				accountDestinationID: 2,
				amount:               50,
			},
			runBefore: func(args args, store store.AccountStore) {
				accountDestination := &entities.Account{
					ID:      2,
					Balance: 200,
				}

				store.Save(accountDestination)
			},

			wantErr: erros.ErrOriginAccNotFound,
		},
		{
			name: "should return ErrInsufficientFunds when account origin balance is lower than transfer amount",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountOriginID:      1,
				accountDestinationID: 2,
				amount:               500,
			},
			runBefore: func(args args, store store.AccountStore) {
				accountOriginTest := &entities.Account{
					ID:      1,
					Balance: 200,
				}
				accountDestinationTest := &entities.Account{
					ID:      2,
					Balance: 300,
				}
				store.Save(accountOriginTest)
				store.Save(accountDestinationTest)
			},

			wantErr: erros.ErrInsufficientFunds,
		},
		{
			name: "should return ErrInvalidAmount when trying to transfer 0 coins",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				accountOriginID:      1,
				accountDestinationID: 2,
				amount:               0,
			},
			runBefore: func(args args, store store.AccountStore) {
				accountOriginTest := &entities.Account{
					ID:      1,
					Balance: 200,
				}
				accountDestinationTest := &entities.Account{
					ID:      2,
					Balance: 300,
				}
				store.Save(accountOriginTest)
				store.Save(accountDestinationTest)
			},
			wantErr: erros.ErrInvalidAmount,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.store.ResetAccountStates()
			if tt.runBefore != nil {
				tt.runBefore(tt.args, tt.fields.store)
			}
			useCase := accounts.NewAccountUseCase(tt.fields.store)

			originAccountGot, destinationAccountGot, err := useCase.Transfer(tt.args.accountOriginID, tt.args.accountDestinationID, tt.args.amount)
			if tt.wantErr != nil {
				if tt.wantErr == genericErr {
					err = genericErr
				}
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				tt.want[0] = originAccountGot
				tt.want[1] = destinationAccountGot
				assert.Equal(t, tt.want[0], originAccountGot)
				assert.Equal(t, tt.want[1], destinationAccountGot)
			}
		},
		)
	}

}
