package main

import (
	"gobackend/dao"
	"gobackend/router"
	"log"
	"net/http"
)

func main() {
	err := dao.ConnectDb("db")
	if err != nil {
		log.Fatal(err)
	}
	r := router.SetupRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
