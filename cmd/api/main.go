package main

import (
	"digital_signage_api/internal"
	"digital_signage_api/internal/ws"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// hanya load .env kalau APP_ENV != production
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Println("No .env file found, using system env")
		}
	}

	// jalankan scheduler websocket (cek tiap menit)
	go ws.RunScheduler()

	// start aplikasi utama (GIN, DB, route, dll.)
	internal.InitApp()
}
