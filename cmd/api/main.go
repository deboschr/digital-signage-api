package main

import (
    "digital_signage_api/internal"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
    // hanya load .env kalau APP_ENV != production
    if os.Getenv("APP_ENV") != "production" {
        if err := godotenv.Load(".env"); err != nil {
            log.Println("No .env file found, using system env")
        }
    }

    internal.InitApp()
}
