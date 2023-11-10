package main

import (
	"backend/api"
	db "backend/db/sqlc"
	"backend/utils"
	"context"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	var err error

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	querier := db.New(connPool)

	server := api.NewServer(querier)

	err = server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
