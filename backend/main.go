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

	println(viper.GetString("DB_SOURCE"))

	//aws_session := session.Must(session.NewSession())
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	//connPool, err := pgxpool.New(context.Background(), viper.Get("DB_SOURCE").(string))
	connPool, err := pgxpool.New(context.Background(), "postgresql://root:secret@db:5432/xtracker?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	queries := db.New(connPool)
	server := api.NewServer(cfg, queries)

	err = server.Start(":8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
