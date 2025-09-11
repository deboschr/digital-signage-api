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
   db.DB.AutoMigrate(
        &models.User{},
        &models.Airport{},
        &models.Device{},
        &models.Channel{},
        &models.Content{},
        &models.Schedule{},
    )

	// init Gin
	r := gin.Default()

	// health endpoints
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/ready", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "degraded"})
			return
		}
		c.JSON(200, gin.H{"status": "ready"})
	})

    // kumpulin semua route entity di sini
	// group API v1
	api := r.Group("/api/v1")
	{
		routes.RegisterDeviceRoutes(api)
		// routes.RegisterUserRoutes(api)
		// routes.RegisterAirportRoutes(api)
		// routes.RegisterChannelRoutes(api)
	}

	// run server
	r.Run(":8080")
}
