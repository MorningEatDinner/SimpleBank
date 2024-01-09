package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/xiaorui/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// 创建一个新的server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//添加一些路由, 如果你传入多个handle函数， 那么最后一个才是处理这个请求的函数， 其他的是中间件
	router.POST("/accounts", server.createAccount) // createAccount这个方法必须访问store才能存储数据
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
