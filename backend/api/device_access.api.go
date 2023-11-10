package api

import (
	db "backend/db/sqlc"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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
	_, err := s.queries.GetUserByID(ctx, req.DeviceID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusFailedDependency, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//check if user exists
	_, err = s.queries.GetDevice(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusFailedDependency, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// todo
	// check if access already exists
	existRows, err := s.queries.GetAccessWithDeviceIDAndUserID(ctx, db.GetAccessWithDeviceIDAndUserIDParams{
		DeviceID: req.DeviceID,
		UserID:   req.UserID,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	if len(existRows) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "access already exists",
		})

		return
	}

	arg := db.CreateAccessParams{
		DeviceID:   req.DeviceID,
		UserID:     req.UserID,
		Permission: req.Permission,
	}

	deviceAccess, err := s.queries.CreateAccess(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//
	ctx.JSON(http.StatusOK, deviceAccess)
}

//func (s *Server) checkDuplicateAccess(ctx *gin.Context, deviceID int64, userID int64) (exists bool, err error) {
//	_, err = s.queries.GetAccessWithDeviceIDAndUserID(ctx, db.GetAccessWithDeviceIDAndUserIDParams{
//		DeviceID: deviceID,
//		UserID:   userID,
//	})
//
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return false, nil
//		}
//
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return false, err
//	}
//
//	ctx.JSON(http.StatusForbidden, errorResponse(err))
//	return true, nil
//}
