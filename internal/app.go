package internal

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/routes"
)

func InitApp() {
    db.Init()
    db.DB.AutoMigrate(&models.Device{})

    r := routes.SetupRouter()
    r.Run(":8080")
}
