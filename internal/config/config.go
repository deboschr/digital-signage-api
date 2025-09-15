package config

import "os"

type Config struct {
    AppPort string
    Env     string

    DBHost string
    DBPort string
    DBName string
    DBUser string
    DBPass string
}

func Load() Config {
    return Config{
        AppPort: getenv("APP_PORT", "8080"),
        Env:     getenv("APP_ENV", "development"),
        DBHost:  getenv("DB_HOST", ""),
        DBPort:  getenv("DB_PORT", "3306"),
        DBName:  getenv("DB_NAME", ""),
        DBUser:  getenv("DB_USER", ""),
        DBPass:  getenv("DB_PASS", ""),
    }
}

func getenv(k, def string) string {
    if v := os.Getenv(k); v != "" {
        return v
    }
    return def
}
