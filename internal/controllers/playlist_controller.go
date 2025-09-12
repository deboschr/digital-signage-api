package controllers

import (
	"digital_signage_api/internal/dto"
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

// GET /playlists
func (c *PlaylistController) GetPlaylists(ctx *gin.Context) {
	playlists, err := c.service.GetAllPlaylists()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// playlists = []dto.SummaryPlaylistDTO
	ctx.JSON(http.StatusOK, playlists)
}

// GET /playlists/:id
func (c *PlaylistController) GetPlaylist(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	playlist, err := c.service.GetPlaylistByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
		return
	}
	// playlist = dto.DetailPlaylistDTO
	ctx.JSON(http.StatusOK, playlist)
}

// POST /playlists
func (c *PlaylistController) CreatePlaylist(ctx *gin.Context) {
	var req dto.CreatePlaylistReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.CreatePlaylist(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// res = dto.CreatePlaylistResDTO
	ctx.JSON(http.StatusCreated, res)
}

// PUT/PATCH /playlists/:id
func (c *PlaylistController) UpdatePlaylist(ctx *gin.Context) {
	var req dto.UpdatePlaylistReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.UpdatePlaylist(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// res = dto.UpdatePlaylistResDTO
	ctx.JSON(http.StatusOK, res)
}

// DELETE /playlists/:id
func (c *PlaylistController) DeletePlaylist(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeletePlaylist(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "playlist deleted"})
}

// -----------------------------
// Playlist Content management
// -----------------------------

// POST /playlists/content
func (c *PlaylistController) CreatePlaylistContent(ctx *gin.Context) {
	var req dto.CreatePlaylistContentReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist, err := c.service.AddContents(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, playlist)
}



// PATCH /playlists/content
func (c *PlaylistController) UpdatePlaylistContent(ctx *gin.Context) {
	var req dto.UpdatePlaylistContentReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist, err := c.service.UpdateOrders(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// playlist = dto.DetailPlaylistDTO
	ctx.JSON(http.StatusOK, playlist)
}

// DELETE /playlists/content
func (c *PlaylistController) DeletePlaylistContent(ctx *gin.Context) {
	var req dto.DeletePlaylistContentReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.RemoveContents(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "contents removed"})
}
