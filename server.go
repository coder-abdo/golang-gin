package main

import (
	"github.com/coder-abdo/gin-tutorial/config"
	"github.com/coder-abdo/gin-tutorial/controller"
	"github.com/coder-abdo/gin-tutorial/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()
	server.Use(gin.Recovery(), middlewares.Logger())
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)

	}
	server.Run(":8080")
}
