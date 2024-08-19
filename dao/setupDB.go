package dao

import (
	"gobackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(dbName string) error {
	database, err := gorm.Open(sqlite.Open(dbName+"test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = database.AutoMigrate(&models.Book{})
	if err != nil {
		return err
	}
	DB = database
	return nil
}
