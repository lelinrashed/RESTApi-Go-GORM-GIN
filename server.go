package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lelinrashed/RESTApi/config"
	"github.com/lelinrashed/RESTApi/controller"
	"github.com/lelinrashed/RESTApi/repository"
	"github.com/lelinrashed/RESTApi/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(userService, jwtService)
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
