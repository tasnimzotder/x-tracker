package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"github.com/tasnimzotder/x-tracker/api"
	db "github.com/tasnimzotder/x-tracker/db/sqlc"
	"github.com/tasnimzotder/x-tracker/utils"
	"log"
)

func main() {
	var err error

	_, err = utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	println(viper.GetString("SERVER_ADDRESS"))

	//aws_session := session.Must(session.NewSession())
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	connPool, err := pgxpool.New(context.Background(), viper.GetString("DB_SOURCE"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	queries := db.New(connPool)
	server := api.NewServer(cfg, queries)

	err = server.Start(viper.GetString("SERVER_ADDRESS"))
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
