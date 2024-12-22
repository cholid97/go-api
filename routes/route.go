package routes

import (
	"github.com/cholid97/go-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", userController.GetUsers)
		userRoutes.GET("/:id", userController.GetUser)
		userRoutes.POST("", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
