package controllers

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/services"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ContentController struct {
	service services.ContentService
}

func NewContentController(service services.ContentService) *ContentController {
	return &ContentController{service}
}

// GET /contents
func (c *ContentController) GetContents(ctx *gin.Context) {
	contents, err := c.service.GetAllContents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// contents = []dto.SummaryContentDTO
	ctx.JSON(http.StatusOK, contents)
}

// GET /contents/:id
func (c *ContentController) GetContent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	content, err := c.service.GetContentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "content not found"})
		return
	}
	// content = dto.DetailContentDTO
	ctx.JSON(http.StatusOK, content)
}

// POST /contents
// multipart/form-data dengan field: file (required)
func (c *ContentController) CreateContent(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// Simpan file ke folder contents/
	savePath := filepath.Join("contents", file.Filename)
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ekst
