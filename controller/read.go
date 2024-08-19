package controller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gobackend/dao"
	"gobackend/models"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	err := dao.GetAll(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errors.New("not found id")
		return
	}

	var book models.Book
	err = dao.GetById(&book, uint(id))
	if err != nil {
		errors.New("not found id ka json")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
