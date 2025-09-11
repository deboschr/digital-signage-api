package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContentRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewContentRepository(db)
	service := services.NewContentService(repo)
	controller := controllers.NewContentController(service)

	content := r.Group("/content")
	{
		content.GET("", controller.GetContents)
		content.GET("/:id", controller.GetContent)
		content.POST("", controller.CreateContent)
		content.PATCH("", controller.UpdateContent)
		content.DELETE("/:id", controller.DeleteContent)
	}
}
