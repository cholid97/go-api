package main

import (
	"github.com/cholid97/go-api/config"
	"github.com/cholid97/go-api/controllers"
	"github.com/cholid97/go-api/repositories"
	"github.com/cholid97/go-api/routes"
	"github.com/cholid97/go-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Setup Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, userController)

	// Start the server
	router.Run(":8080")
}
