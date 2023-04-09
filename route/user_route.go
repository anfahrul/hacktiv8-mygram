package route

import (
	"github.com/anfahrul/hacktiv8-mygram/controller"
	"github.com/anfahrul/hacktiv8-mygram/repository"
	"github.com/anfahrul/hacktiv8-mygram/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SetupUserRoute(router *gin.Engine, db *gorm.DB) {
	validate := validator.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
}
