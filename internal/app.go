package internal

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/routes"
	"digital_signage_api/internal/tcp"
	"os"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApp() {
	// init DB
	db.Init()

	// migrasi database
	// if err := db.DB.AutoMigrate(
	// 	&models.Airport{},
	// 	&models.User{},
	// 	&models.Device{},
	// 	&models.Playlist{},
	// 	&models.Content{},
	// 	&models.PlaylistContent{},
	// 	&models.Schedule{},
	// ); err != nil {
	// 	panic("failed to migrate database: " + err.Error())
	// }

	
	
	// init Gin
	r := gin.Default()

	// aktifkan CORS
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173"}, // asal frontend Vite
   //  	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
   //  	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
   //  	ExposeHeaders:    []string{"Content-Length"},
   //  	AllowCredentials: true,
	// }))
	
	// aktifkan CORS untuk semua origin (development only)
   r.Use(cors.Default())


	
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
	
	// folder statis untuk menyimpan konten
	staticPath := os.Getenv("STATIC_PATH") // "../media"
	r.Static(staticPath, ".."+staticPath)


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

	// TCP
	go tcp.StartTCPServer()


	// run server
	if err := r.Run(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
