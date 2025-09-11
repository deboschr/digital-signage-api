package internal

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApp() {
	// init DB
	db.Init()

	// migrasi database
	if err := db.DB.AutoMigrate(
		&models.Airport{},
		&models.User{},
		&models.Device{},
		&models.Playlist{},
		&models.Content{},
		&models.PlaylistContent{},
		&models.Schedule{},
	); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	// init Gin
	r := gin.Default()

	// health endpoints
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/ready", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "degraded"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	// group API v1
	api := r.Group("/api/v1")
	{
		routes.UserRoutes(api, db.DB)
		routes.AirportRoutes(api, db.DB)
		routes.DeviceRoutes(api, db.DB)
		routes.PlaylistRoutes(api, db.DB)
		routes.ContentRoutes(api, db.DB)
		routes.ScheduleRoutes(api, db.DB)
	}

	// run server
	if err := r.Run(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
