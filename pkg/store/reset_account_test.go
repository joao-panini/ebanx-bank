package store_test

import (
	"testing"

	"github.com/joao-panini/ebanx-bank/pkg/entities"
	"github.com/joao-panini/ebanx-bank/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestStoreReset(t *testing.T) {

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
			name: "Reset account sucess",
			fields: fields{
				store: store.NewAccountStore(),
			},
			runBefore: func(store store.AccountStore) {
				acc := &entities.Account{ID: 1, Balance: 100}
				store.Save(acc)
			},
			wantErr: nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.runBefore != nil {
				tt.runBefore(tt.fields.store)
			}

			err := tt.fields.store.ResetAccountStates()
			if tt.wantErr != nil {
				assert.Contains(t, err, tt.wantErr)
			}
		},
		)
	}

}
