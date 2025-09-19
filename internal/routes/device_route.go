package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/middlewares"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeviceRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewDeviceRepository(db)
	service := services.NewDeviceService(repo)
	controller := controllers.NewDeviceController(service)

	device := r.Group("/device")
	device.Use(middlewares.AuthRequired())
	{
		device.GET("", controller.GetDevices)
		device.GET("/:id", controller.GetDevice)
		device.POST("", controller.CreateDevice)
		device.PATCH("", controller.UpdateDevice)
		device.DELETE("/:id", controller.DeleteDevice)
	}

	r.GET("/device/connect", controller.ConnectDeviceWS)
}
