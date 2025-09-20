package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/middlewares"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	auth := r.Group("/auth")
	{
		auth.GET("verify", controller.Verify)
		auth.POST("signin", controller.SignIn)
		auth.DELETE("signout", middlewares.AuthRequired(), controller.SignOut)
	}

	user := r.Group("/user")
	user.Use(middlewares.AuthRequired())
	{
		user.GET("", controller.GetUsers)
		user.GET("/:id", controller.GetUser)
		user.POST("", controller.CreateUser)
		user.PATCH("", controller.UpdateUser)
		user.DELETE("/:id", controller.DeleteUser)
	}
}
