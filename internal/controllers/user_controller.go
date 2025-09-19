package controllers

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

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

	// === Simpan user_id ke session ===
	session := sessions.Default(ctx)
	userJSON, _ := json.Marshal(user)
	session.Set("user", string(userJSON))
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save session"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "signin success",
		"user":    user,
	})
}


func (c *UserController) SignOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"message": "signout success"})
}


// --- CRUD User ---

func (c *UserController) GetUsers(ctx *gin.Context) {

	users, err := c.service.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := c.service.GetUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

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

	ctx.JSON(http.StatusCreated, res)
}

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

	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	
	if err := c.service.DeleteUser(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
