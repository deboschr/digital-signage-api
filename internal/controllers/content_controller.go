package controllers

import (
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContentController struct {
	service services.ContentService
}

func NewContentController(service services.ContentService) *ContentController {
	return &ContentController{service}
}

func (c *ContentController) GetContents(ctx *gin.Context) {
	contents, err := c.service.GetAllContents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, contents)
}

func (c *ContentController) GetContent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	content, err := c.service.GetContentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "content not found"})
		return
	}
	ctx.JSON(http.StatusOK, content)
}

func (c *ContentController) CreateContent(ctx *gin.Context) {
	var content models.Content
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateContent(&content); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, content)
}

func (c *ContentController) UpdateContent(ctx *gin.Context) {
	var content models.Content
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.UpdateContent(&content); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, content)
}

func (c *ContentController) DeleteContent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteContent(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "content deleted"})
}
