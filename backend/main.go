package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	aws_config "github.com/aws/aws-sdk-go-v2/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/tasnimzotder/x-tracker/api"
	db "github.com/tasnimzotder/x-tracker/db/sqlc"
	"github.com/tasnimzotder/x-tracker/utils"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Printf("No .env file found, falling back to environment variables")
	}
}

func NewTLSConfig() (config *tls.Config, err error) {
	certsPath := utils.GetEnvVariable("CERTS_PATH")

	certPool := x509.NewCertPool()
	//permCerts, err := ioutil.ReadFile("certs/AmazonRootCA1.pem")
	permCerts, err := ioutil.ReadFile(certsPath + "/AmazonRootCA1.pem")
	if err != nil {
		log.Fatal(err)
	}

	certPool.AppendCertsFromPEM(permCerts)

	//cert, err := tls.LoadX509KeyPair("certs/certificate.pem.crt", "certs/private.pem.key")
	cert, err := tls.LoadX509KeyPair(certsPath+"/certificate.pem.crt", certsPath+"/private.pem.key")
	if err != nil {
		log.Fatal(err)
	}

	config = &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{cert},
		// ClientCAs:    nil,
		// ClientAuth:   tls.NoClientCert,
	}

	return
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	var err error

	cfg, err := aws_config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// redis client
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     utils.GetEnvVariable("REDIS_SOURCE"),
	// 	DB:       0,
	// 	Password: utils.GetEnvVariable("REDIS_PASSWORD"),
	// })

	connPool, err := pgxpool.New(context.Background(), utils.GetEnvVariable("DB_SOURCE"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// mqtt

	tlsConfig, err := NewTLSConfig()
	if err != nil {
		log.Fatalf("failed to create tls config: %v", err)
	}

	mqttOpts := mqtt.NewClientOptions()
	mqttOpts.AddBroker(fmt.Sprintf("tls://%s:%s", utils.GetEnvVariable("MQTT_ENDPOINT"), utils.GetEnvVariable("MQTT_PORT")))
	mqttOpts.SetClientID(utils.GetEnvVariable("MQTT_CLIENT_ID")).SetTLSConfig(tlsConfig)
	mqttOpts.SetDefaultPublishHandler(f)
	mqttOpts.SetKeepAlive(30 * time.Second)

	mqttClient := mqtt.NewClient(mqttOpts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to connect to mqtt broker: %v", token.Error())
	}

	// influxdb
	influxdbToken := utils.GetEnvVariable("INFLUXDB_TOKEN")
	influxdbURL := utils.GetEnvVariable("INFLUXDB_URL")
	influxdbClient := influxdb2.NewClient(influxdbURL, influxdbToken)

	// db

	queries := db.New(connPool)
	server := api.NewServer(cfg, queries, mqttClient, influxdbClient)

	err = server.Start(
		utils.GetEnvVariable("SERVER_ADDRESS"),
	)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
