package middlewares

import (
	"encoding/json"
	"net/http"

	"digital_signage_api/internal/dto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authorization(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userData := session.Get("user")
		if userData == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		// decode JSON ke DTO
		var user dto.GetSummaryUserResDTO
		if err := json.Unmarshal([]byte(userData.(string)), &user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid session data"})
			ctx.Abort()
			return
		}

		// cek role user apakah masuk ke dalam daftar roles yang diizinkan
		allowed := false
		for _, role := range roles {
			if user.Role == role {
				allowed = true
				break
			}
		}
		if !allowed {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			ctx.Abort()
			return
		}

		// simpan user ke context agar bisa dipakai di handler berikutnya
		ctx.Set("user", user)

		ctx.Next()
	}
}
