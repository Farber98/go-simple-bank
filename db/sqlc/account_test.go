package db_test

import (
	"context"
	db "go-simple-bank/db/sqlc"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccounts(t *testing.T) {

	t.Run("Create account", func(t *testing.T) {
		args := &db.CreateAccountParams{
			Owner:    "Tomito",
			Balance:  50,
			Currency: "ARS",
		}

		account, err := testQueries.CreateAccount(context.Background(), *args)
		require.NoError(t, err)
		require.NotEmpty(t, account)
		require.Equal(t, args.Owner, account.Owner)
		require.Equal(t, args.Balance, account.Balance)
		require.Equal(t, args.Currency, account.Currency)

		// Check that postgres generates correct values.
		require.NotZero(t, account.ID)
		require.NotZero(t, account.CreatedAt)
	})

	t.Run("Delete Account", func(t *testing.T) {

	})

	t.Run("Get Account", func(t *testing.T) {

	})

	t.Run("Get Account", func(t *testing.T) {

	})

	t.Run("List Accounts", func(t *testing.T) {

	})

	t.Run("Update Account", func(t *testing.T) {

	})
}
