package utils

import (
	"fmt"
	"net/url"
	"os"
)

// BuildContentURL generates a public URL for a content file
func BuildContentURL(title string) string {
	baseHost := os.Getenv("APP_HOST")
	if baseHost == "" {
		baseHost = "localhost"
	}

	basePort := os.Getenv("APP_PORT")
	if basePort == "" {
		basePort = "8080"
	}

	staticPath := os.Getenv("STATIC_PATH")
	if staticPath == "" {
		staticPath = "/contents"
	}

	filename := url.PathEscape(title)
	return fmt.Sprintf("http://%s:%s%s/%s", baseHost, basePort, staticPath, filename)
}
