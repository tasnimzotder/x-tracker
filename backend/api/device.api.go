package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	db "github.com/tasnimzotder/x-tracker/db/sqlc"
	"github.com/tasnimzotder/x-tracker/utils"
	"net/http"
	"time"
)

type createDeviceRequest struct {
	DeviceName string `json:"deviceName" binding:"required"`
	UserID     int    `json:"userID" binding:"required"`
}

func (s *Server) createDevice(ctx *gin.Context) {
	var req createDeviceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.UserID <= 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(
			errors.New("user id must be greater than zero"),
		))
		return
	}

	//check if user exists
	_, err := s.queries.GetUser(ctx, int64(req.UserID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(
			errors.New("user not found"),
		))
		return
	}

	arg := db.CreateDeviceParams{
		DeviceName: req.DeviceName,
		UserID:     int64(req.UserID),
		DeviceKey:  utils.GenerateUUID(),
		Status:     "active",
		CreatedAt:  time.Now(),
	}

	device, err := s.queries.CreateDevice(ctx, arg)
	if err != nil {
		var pqErr *pgconn.PgError

		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, device)
}

func (s *Server) getDeviceByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(
			errors.New("user id is required"),
		),
		)
		return
	}

	parsedUserID, err := utils.ParseInt(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(
			errors.New("user id must be a number"),
		))
		return
	}

	devices, err := s.queries.GetDevicesByUser(ctx, int64(parsedUserID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, devices)
}
