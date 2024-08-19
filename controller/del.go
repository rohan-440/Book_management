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

func Del(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errors.New("id not found")
		return
	}
	var book models.Book
	err = dao.Del(&book, uint(id))
	if err != nil {
		errors.New("book not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Book deleted")

}
