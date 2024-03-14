package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tasnimzotder/x-tracker/db/sqlc"
)

type CreateGeofenceRequest struct {
	DeviceID     int64    `json:"device_id" binding:"required"`
	GeofenceName string   `json:"geofence_name" binding:"required"`
	CenterLat    *float64 `json:"center_lat" binding:"required"`
	CenterLong   *float64 `json:"center_long" binding:"required"`
	Radius       float64  `json:"radius" binding:"required"` // in meters
	Rule         string   `json:"rule" binding:"required"`
}

func (s *Server) createGeofence(ctx *gin.Context) {
	var request CreateGeofenceRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check if device exists
	_, err := s.Queries.GetDevice(ctx, request.DeviceID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(
			errors.New("device not found"),
		))
		return
	}

	// create geofence
	arg := db.CreateGeofenceParams{
		DeviceID:     request.DeviceID,
		GeofenceName: request.GeofenceName,
		CenterLat:    *request.CenterLat,
		CenterLong:   *request.CenterLong,
		Radius:       request.Radius,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Status:       "active",
		Rule:         request.Rule,
	}

	geofence, err := s.Queries.CreateGeofence(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, geofence)
}
