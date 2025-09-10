package controllers

import (
    "digital_signage_api/internal/models"
    "digital_signage_api/internal/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type DeviceController struct {
    service *services.DeviceService
}

func NewDeviceController(service *services.DeviceService) *DeviceController {
    return &DeviceController{service: service}
}

func (c *DeviceController) GetDevices(ctx *gin.Context) {
    devices, err := c.service.GetDevices()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, devices)
}

func (c *DeviceController) CreateDevice(ctx *gin.Context) {
    var device models.Device
    if err := ctx.ShouldBindJSON(&device); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.AddDevice(&device); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, device)
}
