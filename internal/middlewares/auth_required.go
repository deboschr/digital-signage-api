package middlewares

import (
	"encoding/json"
	"net/http"

	"digital_signage_api/internal/dto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
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

		// simpan user ke context agar bisa dipakai di handler berikutnya
		ctx.Set("user", user)

		ctx.Next()
	}
}
