package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() (*gorm.DB, error) {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error)
	}

	if err = db.AutoMigrate(&Grocery{}); err != nil {
		log.Println(err)
	}

	return db, err
}

func GetDB() *gorm.DB {
	return DB
}
