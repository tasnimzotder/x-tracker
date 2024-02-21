package api

import (
	"github.com/IBM/sarama"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	db "github.com/tasnimzotder/x-tracker/db/sqlc"
)

type DynamoDB struct {
	DynamoDBClient *dynamodb.Client
	TableName      string
}

type Server struct {
	Router    *gin.Engine
	AwsConfig aws.Config
	queries   *db.Queries
	// rdb        *redis.Client
	DynamoDB       *DynamoDB
	mqttClient     mqtt.Client
	influxdbClient influxdb2.Client
	kafkaProducer  sarama.SyncProducer
}

func NewServer(
	cfg aws.Config,
	queries *db.Queries,
	mqttClient mqtt.Client,
	influxdbClient influxdb2.Client,
	kafkaProducer sarama.SyncProducer,
) *Server {
	server := &Server{
		AwsConfig: cfg,
		queries:   queries,
		// rdb:        rdb,
		mqttClient:     mqttClient,
		influxdbClient: influxdbClient,
		kafkaProducer:  kafkaProducer,
		DynamoDB: &DynamoDB{
			DynamoDBClient: dynamodb.NewFromConfig(cfg),
			TableName:      "xt_test_edge_data",
		},
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// user
	router.POST("/v1/users/create", server.createUser)
	router.GET("/v1/users/id/:id", server.getUserByID)
	router.GET("/v1/users/username/:username", server.getUserByUsername)
	router.GET("/v1/users/all/:limit/:offset", server.getAllUsers)

	router.PUT("/v1/users/update", server.updateUser)

	router.POST("/v1/users/login", server.userLogin)

	// device
	router.POST("/v1/devices/create", server.createDevice)
	router.GET("/v1/devices/user/:user_id", server.getDeviceByUserID)

	//locations
	//router.POST("/v1/locations/get", server.getLastLocations)
	// router.GET("/v1/locations/user/:user_id/:limit", server.getLocationsByUserID)

	// websocket
	router.GET("/v1/ws/location", server.wsLatestLocation)
	router.POST("/v1/geofence/create", server.createGeofence)

	// test
	//router.POST("/v1/test", server.sendMessagehandler)

	server.Router = router
	return server
}

func (s *Server) Start(address string) error {
	// subscribe to mqtt topic
	if token := s.mqttClient.Subscribe("xt/core/data", 1, s.msgHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// s.WriteDataToInfluxDB()
	//s.setupKafka()

	return s.Router.Run(address)
}

func (s *Server) msgHandler(_ mqtt.Client, msg mqtt.Message) {
	// log.Printf("TOPIC: %s\n", msg.Topic())
	// log.Printf("MSG: %s\n", msg.Payload())

	s.WriteDataToInfluxDB(msg.Payload())
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
