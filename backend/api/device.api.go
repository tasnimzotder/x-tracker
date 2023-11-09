package api

import (
	db "backend/db/sqlc"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createDeviceRequest struct {
	DeviceName string `json:"device_name" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

func (s *Server) createDevice(ctx *gin.Context) {
	var req createDeviceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateDeviceParams{
		DeviceName: req.DeviceName,
		Status:     req.Status,
		CreatedAt:  time.Now(),
	}

	device, err := s.querier.CreateDevice(ctx, arg)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//rsp := createDeviceResponse{
	//	DeviceName: device.DeviceName,
	//	Status:     device.Status,
	//	CreatedAt:  device.CreatedAt,
	//	ID:         device.ID,
	//}

	ctx.JSON(http.StatusOK, device)
}

type getDeviceByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getDeviceByID(ctx *gin.Context) {
	var req getDeviceByIDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	device, err := s.querier.GetDevice(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, device)
}
