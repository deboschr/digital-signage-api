package config

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
   	AllowOrigins:     []string{"https://cms.pivods.com", "http://localhost:5173"},
    	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
    	ExposeHeaders:    []string{"Content-Length"},
    	AllowCredentials: true,
    	MaxAge:           12 * time.Hour,
    	AllowWebSockets:  true,
	}))

}
