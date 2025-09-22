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
	{
		schedule.GET("", middlewares.Authorization("admin", "operator", "management"), controller.GetSchedules)
		schedule.GET("/:id", middlewares.Authorization("admin", "operator", "management"), controller.GetSchedule)
		schedule.POST("", middlewares.Authorization("admin", "operator"), controller.CreateSchedule)
		schedule.PATCH("", middlewares.Authorization("admin", "operator"), controller.UpdateSchedule)
		schedule.DELETE("/:id", middlewares.Authorization("admin", "operator"), controller.DeleteSchedule)
	}
}
