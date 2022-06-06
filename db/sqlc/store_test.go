package db_test

import (
	"context"
	db "go-simple-bank/db/sqlc"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {

	t.Run("Transfer TX", func(t *testing.T) {
		store := db.NewStore(testDb)

		accountFrom := createRandomAccount(t)
		accountTo := createRandomAccount(t)

		// run n concurrent transfer transactions
		n := 5
		amount := int64(10)

		// Channel to communicate errors to testing function becuase transfer is inside go routine.
		errs := make(chan error)
		// Channel to receive results.
		results := make(chan db.TransferTxResult)

		for i := 0; i < n; i++ {
			go func() {
				result, err := store.TransferTx(context.Background(), db.TransferTxParams{
					FromAccountId: accountFrom.ID,
					ToAccountId:   accountTo.ID,
					Amount:        amount,
				})
				errs <- err
				results <- result
			}()
		}

		// checks from outside go routine.
		for i := 0; i < n; i++ {
			err := <-errs
			require.NoError(t, err)

			// checks transfer
			results := <-results
			require.NotEmpty(t, results)
			transfer := results.Transfer
			require.Equal(t, transfer.FromAccountID, accountFrom.ID)
			require.Equal(t, transfer.ToAccountID, accountTo.ID)
			require.Equal(t, transfer.Amount, amount)
			require.NotZero(t, transfer.ID)
			require.NotZero(t, transfer.CreatedAt)
			_, err = store.GetTransfer(context.Background(), transfer.ID)
			require.NoError(t, err)

			// checks entries
			require.NotEmpty(t, results.FromEntry)
			fromEntry := results.FromEntry
			require.Equal(t, fromEntry.AccountID, accountFrom.ID)
			require.Equal(t, fromEntry.Amount, -amount)
			require.NotZero(t, fromEntry.ID)
			require.NotZero(t, fromEntry.CreatedAt)
			_, err = store.GetEntry(context.TODO(), fromEntry.ID)
			require.NoError(t, err)

			require.NotEmpty(t, results.ToEntry)
			toEntry := results.ToEntry
			require.Equal(t, toEntry.AccountID, accountTo.ID)
			require.Equal(t, toEntry.Amount, amount)
			require.NotZero(t, toEntry.ID)
			require.NotZero(t, toEntry.CreatedAt)
			_, err = store.GetEntry(context.TODO(), toEntry.ID)
			require.NoError(t, err)

			// check balances

		}
	})

}
