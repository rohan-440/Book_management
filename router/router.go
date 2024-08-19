package router

import (
	"github.com/gorilla/mux"
	"gobackend/controller"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/insert", controller.InsertBook).Methods("POST")
	r.HandleFunc("/getAll", controller.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetById).Methods("GET")
	r.HandleFunc("/delete/{id}", controller.Del).Methods("DELETE")
	r.HandleFunc("/update/{id}", controller.UpdateBook).Methods("PUT")
	return r
}
