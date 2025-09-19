package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/middlewares"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AirportRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewAirportRepository(db)
	service := services.NewAirportService(repo)
	controller := controllers.NewAirportController(service)

	airport := r.Group("/airport")
	airport.Use(middlewares.AuthRequired())
	{
		airport.GET("", controller.GetAirports)
		airport.GET("/:id", controller.GetAirport)
		airport.POST("", controller.CreateAirport)
		airport.PATCH("", controller.UpdateAirport)
		airport.DELETE("/:id", controller.DeleteAirport)
	}
}
