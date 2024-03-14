package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/tasnimzotder/x-tracker/db/sqlc"
	"github.com/tasnimzotder/x-tracker/utils"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Role      string    `json:"role"`
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
		HashedPassword: req.Password, // todo: hash the password
		CreatedAt:      time.Now(),
		Role:           "user",
	}

	user, err := s.Queries.CreateUser(ctx, arg)
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

	rsp := userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (s *Server) getUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("id is required")))
		return
	}

	idInt, err := utils.ParseInt(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.Queries.GetUser(ctx, int64(idInt))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Role:      user.Role,
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *Server) getUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("username is required")))
		return
	}

	user, err := s.Queries.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Role:      user.Role,
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *Server) getAllUsers(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")

	if limit == "" {
		limit = "10"
	}

	if offset == "" {
		offset = "0"
	}

	limitInt, err := utils.ParseInt(limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offsetInt, err := utils.ParseInt(offset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	users, err := s.Queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  int32(limitInt),
		Offset: int32(offsetInt),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := make([]userResponse, 0)
	for _, user := range users {
		res = append(res, userResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			FirstName: user.FirstName.String,
			LastName:  user.LastName.String,
			Role:      user.Role,
		})
	}

	ctx.JSON(http.StatusOK, res)
}

type updateUserRequest struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (s *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userOld, err := s.Queries.GetUser(ctx, req.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		ID:       req.ID,
		Username: req.Username,
		Email:    req.Email,
		FirstName: pgtype.Text{
			String: req.FirstName,
			Valid:  true,
		},
		LastName: pgtype.Text{
			String: req.LastName,
			Valid:  true,
		},
		UpdatedAt:   time.Now(),
		Role:        userOld.Role,
		Status:      userOld.Status,
		PhoneNumber: userOld.PhoneNumber,
		CountryCode: userOld.CountryCode,
		PostalCode:  userOld.PostalCode,
	}

	user, err := s.Queries.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Role:      user.Role,
	}

	ctx.JSON(http.StatusOK, res)
}
