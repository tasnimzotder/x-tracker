package api

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type GeofenceResponse struct {
	Longitude float32 `json:"lng"`
	Latitude  float32 `json:"lat"`
	Radius    float32 `json:"radius"`
	Condition string  `json:"condition"`
	Status    string  `json:"status"`
}

func (s *Server) getGeofenceData(ctx *gin.Context) {
	var responses []GeofenceResponse

	for i := 0; i < 10; i++ {
		responses = append(responses, GeofenceResponse{
			Longitude: 10.0 + rand.Float32(),
			Latitude:  10.0 + rand.Float32(),
			Radius:    100.0,
			Condition: "inside",
			Status:    "active",
		})
	}

	ctx.JSON(http.StatusOK, responses)
}
