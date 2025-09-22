package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/middlewares"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AirportRoutes(r *gin.RouterGroup, database *gorm.DB) {
	repository := repositories.NewAirportRepository(database)
	service := services.NewAirportService(repository)
	controller := controllers.NewAirportController(service)

	airport := r.Group("/airport")
	{
		airport.GET("", middlewares.Authorization("admin", "operator", "management"), controller.GetAirports)
		airport.GET("/:id", middlewares.Authorization("admin", "operator", "management"), controller.GetAirport)
		airport.POST("", middlewares.Authorization("admin"), controller.CreateAirport)
		airport.PATCH("", middlewares.Authorization("admin"), controller.UpdateAirport)
		airport.DELETE("/:id", middlewares.Authorization("admin"), controller.DeleteAirport)
	}
}
