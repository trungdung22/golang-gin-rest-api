package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection(dbURL string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error)
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Println(err)
	}
	DB = db
	return db, err
}

func GetDB() *gorm.DB {
	return DB
}
