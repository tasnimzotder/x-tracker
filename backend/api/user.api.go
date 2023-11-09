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

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.Password,
		CreatedAt:      time.Now(),
	}

	user, err := s.querier.CreateUser(ctx, arg)
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

	rsp := createUserResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		ID:        user.ID,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type getUserByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getUserByIDResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Server) getUserByID(ctx *gin.Context) {
	var req getUserByIDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.querier.GetUserByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := getUserByIDResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}
