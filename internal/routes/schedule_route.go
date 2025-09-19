package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/middlewares"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ScheduleRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewScheduleRepository(db)
	service := services.NewScheduleService(repo)
	controller := controllers.NewScheduleController(service)

	schedule := r.Group("/schedule")
	schedule.Use(middlewares.AuthRequired())
	{
		schedule.GET("", controller.GetSchedules)
		schedule.GET("/:id", controller.GetSchedule)
		schedule.POST("", controller.CreateSchedule)
		schedule.PATCH("", controller.UpdateSchedule)
		schedule.DELETE("/:id", controller.DeleteSchedule)
	}
}
