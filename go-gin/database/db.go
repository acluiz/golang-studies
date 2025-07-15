package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	m "golang-studies/go-gin/models"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDatabase(){
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	
	if err != nil {
		log.Panic("error while connecting with database")
	}

	DB.AutoMigrate(&m.Students)
}
