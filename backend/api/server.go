package api

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/tasnimzotder/x-tracker/db/sqlc"
)

type Server struct {
	Router      *gin.Engine
	AWS_Session *session.Session
	queries     *db.Queries
}

func NewServer(session *session.Session, queries *db.Queries) *Server {
	server := &Server{
		AWS_Session: session,
		queries:     queries,
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

	router.POST("/v1/users/login", server.userLogin)

	// device
	router.POST("/v1/devices/create", server.createDevice)
	router.GET("/v1/devices/user/:user_id", server.getDeviceByUserID)

	//locations
	router.POST("/v1/locations/get", server.getLastLocations)

	server.Router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
