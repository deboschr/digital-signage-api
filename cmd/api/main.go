package main

import (
    "digital_signage_api/internal"
    "github.com/joho/godotenv"
    "log"
)

func main() {
    // coba load dari root project
    if err := godotenv.Load(".env"); err != nil {
        log.Println("No .env file found, using system env")
    }

    internal.InitApp()
}
