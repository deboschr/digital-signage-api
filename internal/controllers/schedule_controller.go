package controllers

import (
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	service services.ScheduleService
}

func NewScheduleController(service services.ScheduleService) *ScheduleController {
	return &ScheduleController{service}
}

func (c *ScheduleController) GetSchedules(ctx *gin.Context) {
	schedules, err := c.service.GetAllSchedules()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, schedules)
}

func (c *ScheduleController) GetSchedule(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	schedule, err := c.service.GetScheduleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "schedule not found"})
		return
	}
	ctx.JSON(http.StatusOK, schedule)
}

func (c *ScheduleController) CreateSchedule(ctx *gin.Context) {
	var schedule models.Schedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateSchedule(&schedule); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, schedule)
}

func (c *ScheduleController) UpdateSchedule(ctx *gin.Context) {
	var schedule models.Schedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.UpdateSchedule(&schedule); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, schedule)
}

func (c *ScheduleController) DeleteSchedule(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteSchedule(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "schedule deleted"})
}
