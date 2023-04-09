package route

import (
	"github.com/anfahrul/hacktiv8-mygram/controller"
	"github.com/anfahrul/hacktiv8-mygram/middleware"
	"github.com/anfahrul/hacktiv8-mygram/repository"
	"github.com/anfahrul/hacktiv8-mygram/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupSocialRoute(router *gin.Engine, db *gorm.DB) {
	socialRepository := repository.NewSocialRepository(db)
	socialService := service.NewSocialService(socialRepository)
	socialController := controller.NewSocialController(socialService)

	authUser := router.Group("/social-media", middleware.AuthMiddleware)
	{
		authUser.POST("", socialController.CreateSocial)
		authUser.GET("", socialController.GetAll)
		authUser.GET("/:social_media_id", socialController.GetOne)
		authUser.PUT("/:social_media_id", socialController.UpdateSocialMedia)
		authUser.DELETE("/:social_media_id", socialController.DeleteSocialMedia)
	}
}
