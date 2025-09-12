package controllers

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

// --- Auth ---
// POST /auth/signin
func (c *UserController) SignIn(ctx *gin.Context) {
	var payload struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.Authenticate(payload.Username, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// TODO: generate JWT nanti
	ctx.JSON(http.StatusOK, gin.H{
		"message": "signin success",
		"user":    user, // bisa SummaryUserDTO atau DetailUserDTO sesuai kebutuhan
	})
}

// POST /auth/signout
func (c *UserController) SignOut(ctx *gin.Context) {
	// TODO: invalidate JWT/session
	ctx.JSON(http.StatusOK, gin.H{"message": "signout success"})
}

// --- CRUD User ---

// GET /users
func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// users = []dto.SummaryUserDTO
	ctx.JSON(http.StatusOK, users)
}

// GET /users/:id
func (c *UserController) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := c.service.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	// user = dto.DetailUserDTO
	ctx.JSON(http.StatusOK, user)
}

// POST /users
func (c *UserController) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// res = dto.CreateUserResDTO
	ctx.JSON(http.StatusCreated, res)
}

// PUT/PATCH /users/:id
func (c *UserController) UpdateUser(ctx *gin.Context) {
	var req dto.UpdateUserReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.UpdateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// res = dto.UpdateUserResDTO
	ctx.JSON(http.StatusOK, res)
}

// DELETE /users/:id
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteUser(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
