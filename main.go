package main

import (
	"log"

	"github.com/anfahrul/hacktiv8-mygram/database"
	"github.com/anfahrul/hacktiv8-mygram/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const PORT = ":8080"

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

	router.Run(PORT)
}
