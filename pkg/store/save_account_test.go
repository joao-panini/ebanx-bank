package store_test

import (
	"testing"

	"github.com/joao-panini/ebanx-bank/pkg/entities"
	"github.com/joao-panini/ebanx-bank/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestStoreSave(t *testing.T) {

	type fields struct {
		store store.AccountStore
	}
	type args struct {
		account *entities.Account
	}
	testCases := []struct {
		name      string
		fields    fields
		args      args
		runBefore func(args args, store store.AccountStore)
		want      *entities.Account
		wantErr   string
	}{
		{
			name: "Save new account sucess",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				account: &entities.Account{
					ID:      1,
					Balance: 100,
				},
			},
			want: &entities.Account{
				ID:      1,
				Balance: 100,
			},
		},
		{
			name: "Save new account with invalid id should return error",
			fields: fields{
				store: store.NewAccountStore(),
			},
			args: args{
				account: &entities.Account{
					ID:      0,
					Balance: 100,
				},
			},

			want:    &entities.Account{},
			wantErr: store.ErrEmptyID.Error(),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.store.ResetAccountStates()
			if tt.runBefore != nil {
				tt.runBefore(tt.args, tt.fields.store)
			}

			accStore := store.NewAccountStore()
			got, err := accStore.Save(tt.args.account)
			if tt.wantErr != "" {
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				tt.want.ID = got.ID
				tt.want.Balance = got.Balance
				assert.Equal(t, tt.want, got)
			}
		},
		)
	}

}
