package route

import (
	"net/http"

	"github.com/anfahrul/hacktiv8-mygram/controller"
	"github.com/anfahrul/hacktiv8-mygram/middleware"
	"github.com/anfahrul/hacktiv8-mygram/repository"
	"github.com/anfahrul/hacktiv8-mygram/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPhotoRoute(router *gin.Engine, db *gorm.DB) {
	photoRepository := repository.NewPhotoRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	photoService := service.NewPhotoService(photoRepository, commentRepository)
	photoController := controller.NewPhotoController(photoService)

	authUser := router.Group("/photos", middleware.AuthMiddleware)
	{
		authUser.POST("", photoController.CreatePhoto)
		authUser.GET("", photoController.GetAll)
		authUser.GET("/:photo_id", photoController.GetOne)
		authUser.PUT("/:photo_id", photoController.UpdatePhoto)
		authUser.DELETE("/:photo_id", photoController.DeletePhoto)
	}
}

func test2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello")
}
