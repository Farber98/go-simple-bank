package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	mockdb "go-simple-bank/db/mock"
	db "go-simple-bank/db/sqlc"
	"go-simple-bank/util"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}
}

func TestAPI(t *testing.T) {

	t.Run("Get Account API", func(t *testing.T) {
		account := randomAccount()

		ctrl := gomock.NewController(t)

		store := mockdb.NewMockStore(ctrl)

		// Build stub
		store.EXPECT().GetAccount(gomock.Any(), account.ID).Times(1).Return(account, nil)

		// start test server and send request
		server := NewServer(store)
		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/accounts/%d", account.ID)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, req)

		//check response
		require.Equal(t, http.StatusOK, recorder.Code)

		//check body
		requireBodyMatchAccount(t, recorder.Body, account)
	})
}

// requireBodyMatchAccount checks that the account returned is equal to the account generated.
func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, gotAccount, account)
}
