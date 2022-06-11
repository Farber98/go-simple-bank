package api

import (
	"database/sql"
	"fmt"
	db "go-simple-bank/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createTransferRequest struct {
	FromAccountId int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountId   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (s *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if !s.validAccount(ctx, req.FromAccountId, req.Currency) {
		return
	}
	if !s.validAccount(ctx, req.ToAccountId, req.Currency) {
		return
	}

	args := db.TransferTxParams{
		FromAccountId: req.FromAccountId,
		ToAccountId:   req.ToAccountId,
		Amount:        req.Amount,
	}

	transferResult, err := s.store.TransferTx(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			err := fmt.Errorf("account [%d] hasn't got %s%d to transfer", req.FromAccountId, req.Currency, req.Amount)
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transferResult)
}

// validAccount checks if an account exists and if the currencies match.
func (s *Server) validAccount(ctx *gin.Context, accountId int64, currency string) bool {
	account, err := s.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountId, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}
