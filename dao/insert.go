package dao

import (
	"errors"
	"gobackend/models"
)

func InsertBook(book *models.Book) error {
	if DB == nil {
		return errors.New("DB is not initialized")
	}
	result := DB.Create(&book)
	return result.Error
}

//func InsertBook(book *models.Book) error {
//	if DB == nil {
//		return errors.New("DB is not initialized")
//	}
//	result := DB.Create(&book)
//	return result.Error
//}
