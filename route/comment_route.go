package route

import (
	"github.com/anfahrul/hacktiv8-mygram/controller"
	"github.com/anfahrul/hacktiv8-mygram/middleware"
	"github.com/anfahrul/hacktiv8-mygram/repository"
	"github.com/anfahrul/hacktiv8-mygram/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCommentRoute(router *gin.Engine, db *gorm.DB) {
	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentController := controller.NewCommentController(commentService)

	authUser := router.Group("/comments", middleware.AuthMiddleware)
	{
		authUser.POST("/:photo_id", commentController.CreateComment)
		authUser.GET("", commentController.GetAll)
		authUser.GET("/:comment_id", commentController.GetOne)
		authUser.PUT("/:comment_id", commentController.UpdateComment)
		authUser.DELETE("/:comment_id", commentController.DeleteComment)
	}
}
