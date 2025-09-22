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
	{
		device.GET("", middlewares.Authorization("admin", "operator", "management"), controller.GetDevices)
		device.GET("/:id", middlewares.Authorization("admin", "operator", "management"), controller.GetDevice)
		device.POST("", middlewares.Authorization("admin", "operator"), controller.CreateDevice)
		device.PATCH("", middlewares.Authorization("admin", "operator"), controller.UpdateDevice)
		device.DELETE("/:id", middlewares.Authorization("admin", "operator"), controller.DeleteDevice)
	}

	r.GET("/device/connect", controller.ConnectDeviceWS)
}
