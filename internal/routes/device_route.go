package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterDeviceRoutes(rg *gin.RouterGroup) {
    // wiring dependency (repo → service → controller)
    repo := repositories.NewDeviceRepository()
    service := services.NewDeviceService(repo)
    controller := controllers.NewDeviceController(service)

    devices := rg.Group("/devices")
    {
        devices.GET("/", controller.GetDevices)
        devices.POST("/", controller.CreateDevice)
    }
}
