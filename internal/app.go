package internal

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/routes"
)

func InitApp() {
    db.Init()

    db.DB.AutoMigrate(
        &models.User{},
        &models.Airport{},
        &models.Device{},
        &models.Channel{},
        &models.Content{},
        &models.Schedule{},
    )


    r := routes.SetupRouter()
    r.Run(":8080")
}
