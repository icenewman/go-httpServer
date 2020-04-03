package main

import (
	"./router"
	"log"
	"net/http"
)

var Port = ":8080"

func main() {
	router := router.Router()
	log.Println("server start with port" + Port)
	log.Fatal(http.ListenAndServe(Port, router))
}
