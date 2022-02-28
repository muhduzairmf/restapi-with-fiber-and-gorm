package database

import (
	"log"
	"os"
	"restapi-with-fiber-and-gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb()  {
	db, err := gorm.Open(sqlite.Open("./database/data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect with database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Successfully connected to the database.")
	db.Logger = logger.Default.LogMode(logger.Info)
	
	log.Println("Running migrations.")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}

}
