package database

import (
	"log"
	"project-4/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := GetDBConfig()

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Debug().AutoMigrate(models.User{}, models.Category{}, models.Product{}, models.TransactionHistory{})
}

func GetDB() *gorm.DB {
	return db
}
