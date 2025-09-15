package utils

import (
	"fmt"
	"net/url"
	"os"
)

// BuildContentURL generates a public URL for a content file
func BuildContentURL(title string) string {
    env := os.Getenv("APP_ENV")
    staticPath := os.Getenv("STATIC_PATH")
    if staticPath == "" {
        staticPath = "/media"
    }

    filename := url.PathEscape(title)

    if env == "production" {
        // production → pakai domain publik
        return fmt.Sprintf("https://cms.pivods.com%s/%s", staticPath, filename)
    }

    // default (development / staging) → pakai host:port dari env
    baseHost := os.Getenv("APP_HOST")
    if baseHost == "" {
        baseHost = "localhost"
    }

    basePort := os.Getenv("APP_PORT")
    if basePort == "" {
        basePort = "8080"
    }

    return fmt.Sprintf("http://%s:%s%s/%s", baseHost, basePort, staticPath, filename)
}
