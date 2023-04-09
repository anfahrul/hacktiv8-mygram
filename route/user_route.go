package route

import (
	"github.com/anfahrul/hacktiv8-mygram/controller"
	"github.com/anfahrul/hacktiv8-mygram/middleware"
	"github.com/anfahrul/hacktiv8-mygram/repository"
	"github.com/anfahrul/hacktiv8-mygram/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	socialRepository := repository.NewSocialRepository(db)
	userService := service.NewUserService(userRepository, photoRepository, socialRepository)
	userController := controller.NewUserController(userService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	authUser := router.Group("/user", middleware.AuthMiddleware)
	{
		authUser.GET("/profile", userController.GetProfile)
	}
}
