package controllers

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/services"
	
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	service services.DeviceService
}

func NewDeviceController(service services.DeviceService) *DeviceController {
	return &DeviceController{service}
}

func (c *DeviceController) GetDevices(ctx *gin.Context) {

	devices, err := c.service.GetDevices()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, devices)
}


func (c *DeviceController) GetDevice(ctx *gin.Context) {
	
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	device, err := c.service.GetDevice(uint(id))
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
		return
	}
	
	ctx.JSON(http.StatusOK, device)
}

func (c *DeviceController) CreateDevice(ctx *gin.Context) {
	
	var req dto.CreateDeviceReqDTO
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.CreateDevice(req)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (c *DeviceController) UpdateDevice(ctx *gin.Context) {
	
	var req dto.UpdateDeviceReqDTO
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.UpdateDevice(req)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DeviceController) DeleteDevice(ctx *gin.Context) {
	
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	if err := c.service.DeleteDevice(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "device deleted"})
}

func (c *DeviceController) ConnectDeviceWS(ctx *gin.Context) {
	apiKey := ctx.Query("api_key")

	if apiKey == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "api_key required"})
		return
	}

	// delegasi ke service untuk validasi & start koneksi
	c.service.ConnectDeviceWS(ctx.Writer, ctx.Request, apiKey)
}