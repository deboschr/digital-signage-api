package routes

import (
	"digital_signage_api/internal/controllers"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	auth := r.Group("/auth")
	{
		auth.POST("/signin", controller.SignIn)
		auth.DELETE("/signout", controller.SignOut)
	}

	user := r.Group("/user")
	{
		user.GET("", controller.GetUsers)
		user.GET("/:id", controller.GetUser)
		user.POST("", controller.CreateUser)
		user.PATCH("", controller.UpdateUser)
		user.DELETE("/:id", controller.DeleteUser)
	}
}
