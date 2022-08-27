package main

import (
	"database/sql"
	"log"

	"github.com/copter01252/simplebank/api"
	db "github.com/copter01252/simplebank/db/sqlc"
	"github.com/copter01252/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.Loadconfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServe(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
