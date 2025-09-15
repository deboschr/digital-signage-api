package config

import (
	"os"
)

type Config struct {
	AppPort string
	Env     string

	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string

	StaticPath string
}

// helper ambil env dengan default
func getenv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Load config dari environment
func Load() *Config {
	return &Config{
		AppPort: getenv("APP_PORT", "8080"),
		Env:     getenv("APP_ENV", "development"),

		DBHost: getenv("DB_HOST", "127.0.0.1"),
		DBPort: getenv("DB_PORT", "3306"),
		DBName: getenv("DB_NAME", "digital_signage"),
		DBUser: getenv("DB_USER", "root"),
		DBPass: getenv("DB_PASS", ""),

		// default ./media kalau tidak ada env
		StaticPath: getenv("STATIC_PATH", "./media"),
	}
}
