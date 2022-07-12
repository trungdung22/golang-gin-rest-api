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

	v1 := router.Group("/api")
	endpoints.UsersRegister(v1.Group("/users"))

	v1.Use(middlerwares.AuthMiddleware(false))

	log.Fatal(router.Run(":10000"))
}
