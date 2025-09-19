package controllers

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/services"
	"net/http"
	"os"
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

func (c *ContentController) GetContents(ctx *gin.Context) {
	
	contents, err := c.service.GetContents()
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, contents)
}

func (c *ContentController) GetContent(ctx *gin.Context) {
	
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	content, err := c.service.GetContent(uint(id))
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "content not found"})
		return
	}
	
	ctx.JSON(http.StatusOK, content)
}

func (c *ContentController) CreateContent(ctx *gin.Context) {
	
	file, err := ctx.FormFile("file")
	userVal, _ := ctx.Get("user")
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	userSession := userVal.(dto.GetSummaryUserResDTO)

	// Simpan file ke folder media (STATIC_PATH)
	mediaDir := os.Getenv("STATIC_PATH")
	if mediaDir == "" {
    	mediaDir = "./media"
	}

	savePath := filepath.Join(mediaDir, file.Filename)
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
    	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    	return
	}


	// Ekstrak metadata dari file
	title := filepath.Base(file.Filename)
	ext := strings.ToLower(filepath.Ext(file.Filename))

	ctype := "image"
	if ext == ".mp4" || ext == ".mov" || ext == ".avi" {
		ctype = "video"
	}


	// Default duration = 0 (foto). Bisa dihitung pakai ffprobe untuk video.
	duration := 0

	// Bangun request DTO
	req := dto.CreateContentReqDTO{
		AirportID: *userSession.AirportID,
		Title:    title,
		Type:     ctype,
		Duration: uint16(duration),
	}

	// Service handle mapping ke model + DB
	res, err := c.service.CreateContent(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// res = dto.CreateContentResDTO
	ctx.JSON(http.StatusCreated, res)
}


// DELETE /contents/:id
func (c *ContentController) DeleteContent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteContent(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "content deleted"})
}
