package config

import (
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupSession(r *gin.Engine) {
	// ambil secret dari environment (jangan hardcode)
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "fallback-secret" // default, tapi jangan dipakai di production
	}

	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,             // aktifkan Secure di production (butuh HTTPS)
		MaxAge:   int((24 * time.Hour).Seconds()), // 1 hari
		SameSite: 3,                // SameSite=Lax
	})

	r.Use(sessions.Sessions("cms_session", store))
}
