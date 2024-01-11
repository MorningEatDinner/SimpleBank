package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/xiaorui/simplebank/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currecy       string `json:"currecy" binding:"required,currecy"` //我想应该是汇率在代码中调整吧？就是
	//说最终进行计算的时候都是调整到同一个货币种类下
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // 如果解析数据的时候发生了错误， 比如json传输进来的数据是不满足要求的
		return
	}

	//就是说发起请求有货币类型， 用户账户有货币类型， 他们必须对应上。
	if !server.validAccount(ctx, req.FromAccountID, req.Currecy) {
		//我想这里还要加上ctx.json处理的错误信息？
		//不对， 函数最里面处理的时候已经使用了ctx.json来进行错误处理了， 所以这里不需要
		return
	}
	if !server.validAccount(ctx, req.ToAccountID, req.Currecy) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		//.json就是返回一个response
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currecy string) bool {
	//验证这个账户的货币类型是否是currecy
	account, err := server.store.GetAccount(ctx, accountID)
	//获取账户信息
	if err != nil {
		if err == sql.ErrNoRows {
			//如果没有找到数据
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		//如果发生错误了
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	if account.Currecy != currecy {
		err = fmt.Errorf("account [%d] currecy mismatch: %s vs %s", accountID, account.Currecy, currecy)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
