package main

import (
	"./router"
	"log"
	"net/http"
)

//web server port
var Port = ":80"

func main() {
	log.Println("server start with port" + Port)
	for _, v := range router.Routes{
		log.Printf("Route registy name=%s path=%s method=%s func=%p",v.Name, v.Pattern, v.Method, v.HandlerFunc)
	}
	log.Fatal(http.ListenAndServe(Port, router.MuxRouter))
}
