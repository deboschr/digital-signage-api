package controllers

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AirportController struct {
	service services.AirportService
}

func NewAirportController(service services.AirportService) *AirportController {
	return &AirportController{service}
}

func (c *AirportController) GetAirports(ctx *gin.Context) {

	airports, err := c.service.GetAllAirports()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, airports)
}

func (c *AirportController) GetAirport(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	airport, err := c.service.GetAirportByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "airport not found"})
		return
	}
	
	ctx.JSON(http.StatusOK, airport)
}

func (c *AirportController) CreateAirport(ctx *gin.Context) {

	var req dto.CreateAirportReqDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.CreateAirport(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (c *AirportController) UpdateAirport(ctx *gin.Context) {

	var req dto.UpdateAirportReqDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.UpdateAirport(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *AirportController) DeleteAirport(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	
	if err := c.service.DeleteAirport(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "airport deleted"})
}
