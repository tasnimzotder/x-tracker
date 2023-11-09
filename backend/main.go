package main

import (
	"backend/api"
	db "backend/db/sqlc"
	"backend/utils"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	var err error

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	querier := db.New(conn)

	server := api.NewServer(querier)

	err = server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
