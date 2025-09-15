package utils

import (
	"fmt"
	"net/url"
	"os"
)

func BuildContentURL(title string) string {
	env := os.Getenv("APP_ENV")
	filename := url.PathEscape(title)

	// Kalau production → domain https://cms.pivods.com
	if env == "production" {
		return fmt.Sprintf("https://cms.pivods.com/api/media/%s", filename)
	}

	// Kalau development → tetap pakai host:port dari env
	baseHost := os.Getenv("APP_HOST")
	if baseHost == "" {
		baseHost = "localhost"
	}

	basePort := os.Getenv("APP_PORT")
	if basePort == "" {
		basePort = "8080"
	}

	return fmt.Sprintf("http://%s:%s/media/%s", baseHost, basePort, filename)
}
