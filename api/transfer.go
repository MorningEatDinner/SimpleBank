package api

import (
	db "github.com/xiaorui/simplebank/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required, min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required, min=1"`
	Amount        int64  `json:"amount" binding:"required gt=0"`
	Currecy       string `json:"currecy" binding:"required, oneof=USD EUR CAD"` //我想应该是汇率在代码中调整吧？就是
	//说最终进行计算的时候都是调整到同一个货币种类下
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TransferTxParams{
		Owner:   req.Owner,
		Currecy: req.Currecy,
		Balance: 0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		//.json就是返回一个response
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
