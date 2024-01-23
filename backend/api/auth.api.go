package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type userLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// take userResponse

func (s *Server) userRegister(ctx *gin.Context) {
	s.createUser(ctx)
}

func (s *Server) userLogin(ctx *gin.Context) {
	var req userLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.queries.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// todo: implement
	hashedPassword := req.Password

	if hashedPassword != user.HashedPassword {
		ctx.JSON(http.StatusUnauthorized, errorResponse(
			errors.New("invalid password"),
		))
		return
	}

	rsp := userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Role:      user.Role,
	}

	ctx.JSON(http.StatusOK, rsp)
}
