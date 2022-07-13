package main

import (
	"log"

	"crud-api/config"
	"crud-api/endpoints"
	"crud-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}
	db, err := models.Connection(config.DBUri)
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()

	api := router.Group("/api")
	endpoints.UsersRegisterRouter(api.Group("/users"))

	log.Fatal(router.Run(":8080"))
}
