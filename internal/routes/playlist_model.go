package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PlaylistRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewPlaylistRepository(db)
	service := services.NewPlaylistService(repo)
	controller := controllers.NewPlaylistController(service)

	playlist := r.Group("/playlist")
	{
		playlist.GET("", controller.GetPlaylists)
		playlist.GET("/:id", controller.GetPlaylist)
		playlist.POST("", controller.CreatePlaylist)
		playlist.PATCH("", controller.UpdatePlaylist)
		playlist.DELETE("/:id", controller.DeletePlaylist)
	}
}
