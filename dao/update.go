package dao

import (
	"errors"
	"gobackend/models"
)

func UpdateById(book *models.Book) error {
	if DB == nil {
		return errors.New("Database not initialized")
	}
	result := DB.Save(book)
	return result.Error
}
