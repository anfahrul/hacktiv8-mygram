package main

import (
	"log"

	_ "github.com/anfahrul/hacktiv8-mygram/docs"

	"github.com/anfahrul/hacktiv8-mygram/database"
	"github.com/anfahrul/hacktiv8-mygram/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const PORT = ":8080"

// @title					MyGram API
// @version					1.0
// @description				This is a MyGram API.
// @host 					localhost:8080
// @BasePath 				/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	router := gin.Default()

	database.StartDB()
	db := database.GetDB()

	route.SetupUserRoute(router, db)
	route.SetupPhotoRoute(router, db)
	route.SetupSocialRoute(router, db)
	route.SetupCommentRoute(router, db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}
