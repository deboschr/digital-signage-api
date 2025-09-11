package controllers

import (
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaylistController struct {
	service services.PlaylistService
}

func NewPlaylistController(service services.PlaylistService) *PlaylistController {
	return &PlaylistController{service}
}

func (c *PlaylistController) GetPlaylists(ctx *gin.Context) {
	playlists, err := c.service.GetAllPlaylists()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, playlists)
}

func (c *PlaylistController) GetPlaylist(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	playlist, err := c.service.GetPlaylistByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}

func (c *PlaylistController) CreatePlaylist(ctx *gin.Context) {
	var playlist models.Playlist
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreatePlaylist(&playlist); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, playlist)
}

func (c *PlaylistController) UpdatePlaylist(ctx *gin.Context) {
	var playlist models.Playlist
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.UpdatePlaylist(&playlist); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}

func (c *PlaylistController) DeletePlaylist(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeletePlaylist(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "playlist deleted"})
}


// POST /playlist/content
func (c *PlaylistController) CreatePlaylistContent(ctx *gin.Context) {
	var req struct {
		PlaylistID uint   `json:"PlaylistID" binding:"required"`
		ContentIDs []uint `json:"ContentIDs" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.AddContents(req.PlaylistID, req.ContentIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "contents added"})
}

// PATCH /playlist/content
func (c *PlaylistController) UpdatePlaylistContent(ctx *gin.Context) {
	var req struct {
		PlaylistID uint                    `json:"PlaylistID" binding:"required"`
		Contents   []models.PlaylistContent `json:"Contents" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.UpdateOrders(req.PlaylistID, req.Contents); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "orders updated"})
}

// DELETE /playlist/content
func (c *PlaylistController) DeletePlaylistContent(ctx *gin.Context) {
	var req struct {
		PlaylistID uint   `json:"PlaylistID" binding:"required"`
		ContentIDs []uint `json:"ContentIDs" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.RemoveContents(req.PlaylistID, req.ContentIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "contents removed"})
}
