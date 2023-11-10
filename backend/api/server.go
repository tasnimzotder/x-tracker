package api

import (
	db "backend/db/sqlc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Session *session.Session
	queries *db.Queries
}

func NewServer(queries *db.Queries) *Server {
	server := &Server{
		queries: queries,
	}
	router := gin.Default()

	// aws
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
	}))

	//router.GET("/ping", server.ping)

	router.POST("/v1/locations", server.getLastLocations)
	// get the geofence data
	router.POST("/v1/geofence/get", server.getGeofenceData)
	//// add a geofence
	//router.POST("/geofence/add", server.addGeofence)
	//// delete a geofence
	//router.POST("/geofence/delete", server.deleteGeofence)

	// user
	router.POST("/v1/users", server.createUser)
	router.GET("/v1/users/id/:id", server.getUserByID)

	// device
	router.POST("/v1/devices", server.createDevice)
	router.GET("/v1/devices/id/:id", server.getDeviceByID)

	//// access
	router.POST("/v1/access", server.createDeviceAccess)

	// mqtt listener
	server.mqttListener()

	server.Session = sess
	server.Router = router

	return server
}

func (s *Server) Start(address string) error {
	return s.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
