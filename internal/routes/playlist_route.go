package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/middlewares"
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
		playlist.GET("", middlewares.Authorization("admin", "operator", "management"), controller.GetPlaylists)
		playlist.GET("/:id", middlewares.Authorization("admin", "operator", "management"), controller.GetPlaylist)
		playlist.POST("", middlewares.Authorization("admin", "operator"), controller.CreatePlaylist)
		playlist.PATCH("", middlewares.Authorization("admin", "operator"), controller.UpdatePlaylist)
		playlist.DELETE("/:id", middlewares.Authorization("admin", "operator"), controller.DeletePlaylist)
	}


	content := r.Group("/playlist/content")
	{
		content.POST("", middlewares.Authorization("admin", "operator"), controller.CreatePlaylistContent)
		content.PATCH("", middlewares.Authorization("admin", "operator"), controller.UpdatePlaylistContent)
		content.DELETE("", middlewares.Authorization("admin", "operator"), controller.DeletePlaylistContent)
	}

}
