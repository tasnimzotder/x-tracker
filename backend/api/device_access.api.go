package api

import (
	db "backend/db/sqlc"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createDeviceAccessRequest struct {
	DeviceID   int64  `json:"device_id" binding:"required"`
	UserID     int64  `json:"user_id" binding:"required"`
	Permission string `json:"permission" binding:"required"`
}

func (s *Server) createDeviceAccess(ctx *gin.Context) {
	var req createDeviceAccessRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check if device exists
	_, err := s.querier.GetAccessWithDeviceID(ctx, req.DeviceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusFailedDependency, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// check if user exists
	_, err = s.querier.GetAccessWithUserID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusFailedDependency, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// check if same access already exists with same device_id and user_id
	exists, err := s.checkDuplicateAccess(ctx, req.DeviceID, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if exists {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	arg := db.CreateAccessParams{
		DeviceID:   req.DeviceID,
		UserID:     req.UserID,
		Permission: req.Permission,
	}

	deviceAccess, err := s.querier.CreateAccess(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, deviceAccess)
}

func (s *Server) checkDuplicateAccess(ctx *gin.Context, deviceID int64, userID int64) (exists bool, err error) {
	arg := db.GetAccessWithDeviceIDAndUserIDParams{
		DeviceID: deviceID,
		UserID:   userID,
	}

	deviceAccess, err := s.querier.GetAccessWithDeviceIDAndUserID(ctx, arg)
	if err != nil {
		//if errors.Is(err, sql.ErrNoRows) {
		//	return false, nil
		//}

		return false, err
	}

	if deviceAccess.ID != 0 {
		return true, nil
	}

	return false, nil
}
