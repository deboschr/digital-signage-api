package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/middlewares"
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
		content.GET("", middlewares.Authorization("admin", "operator", "management"), controller.GetContents)
		content.GET("/:id", middlewares.Authorization("admin", "operator", "management"), controller.GetContent)
		content.POST("", middlewares.Authorization("admin", "operator"), controller.CreateContent)
		content.DELETE("/:id", middlewares.Authorization("admin", "operator"), controller.DeleteContent)
	}
}
