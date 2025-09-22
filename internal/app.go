package internal

import (
	"digital_signage_api/internal/config"
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/routes"
	"digital_signage_api/internal/models"

	// "github.com/gin-contrib/cors"
	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	// "os"
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

	
	// setup CORS & Session dari config (production)
	config.SetupCORS(r)
	config.SetupSession(r)
	
	// setup CORS & Session dari config (development)
	// r.Use(cors.Default())
	// store := cookie.NewStore([]byte("super-secret-key"))
	// r.Use(sessions.Sessions("my_session", store))

	cfg := config.Load()
	r.Static("/media", cfg.StaticPath)


	// group API v1
	api := r.Group("/v1")
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
