package dao

import (
	"errors"
	"gobackend/models"
)

func GetAll(books *[]models.Book) error {
	if DB == nil {
		return errors.New("Database not initialized")
	}
	result := DB.Find(&books)
	return result.Error
}
func GetById(book *models.Book, id uint) error {
	if DB == nil {
		return errors.New("Database not initialized")
	}
	result := DB.First(book, id)
	return result.Error
}
