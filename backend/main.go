package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

	//	aws session
	aws_session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatalf("failed to create aws session: %v", err)
	}

	connPool, err := pgxpool.New(context.Background(), viper.GetString("DB_SOURCE"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	queries := db.New(connPool)
	server := api.NewServer(aws_session, queries)

	err = server.Start(viper.GetString("SERVER_ADDRESS"))
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
