package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/xiaorui/simplebank/api"
	db "github.com/xiaorui/simplebank/db/sqlc"
	"github.com/xiaorui/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".") //他会自动去找到这个路径下的app.env
	if err != nil {
		log.Fatal("cannot load config.", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
