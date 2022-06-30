package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lelinrashed/RESTApi/config"
	"github.com/lelinrashed/RESTApi/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	// Auth route group
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	// Run the server
	r.Run()
}
