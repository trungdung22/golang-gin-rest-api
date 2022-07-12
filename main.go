package main

import (
	"log"

	"crud-api/endpoints"
	"crud-api/middlerwares"
	"crud-api/model"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()

	api := router.Group("/api")
	api.Use(middlerwares.UserLoaderMiddleware())
	endpoints.UsersRegisterRouter(api.Group("/users"))

	log.Fatal(router.Run(":10000"))
}
