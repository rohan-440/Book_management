package dao

import (
	"errors"
	"gobackend/models"
)

func Del(book *models.Book, id uint) error {
	if DB == nil {
		return errors.New("database not initialized")
	}
	res := DB.Delete(book, id)
	return res.Error
}
