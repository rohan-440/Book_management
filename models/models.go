package models

type Book struct {
	Id       uint   `json:"id"`
	BookName string `json:"book_name"`
	Author   string `json:"author"`
}
