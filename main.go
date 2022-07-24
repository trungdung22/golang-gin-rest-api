package main

import (
	"fmt"
	"log"
	"os"

	"crud-api/config"
	"crud-api/endpoints"
	"crud-api/middlerwares"
	"crud-api/models"
	"crud-api/seeds"

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

	args := os.Args
	fmt.Println(args)
	if len(args) > 1 {
		first := args[1]
		if first == "seed" {
			seeds.Seed()
			os.Exit(0)
		}
	}

	router := gin.Default()

	api := router.Group("/api")
	api.Use(middlerwares.UserLoaderMiddleware())
	endpoints.UsersRegisterRouter(api.Group("/users"))

	log.Fatal(router.Run(":8080"))
}
