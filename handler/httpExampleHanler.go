package handler

import (
	"../service"
	"encoding/json"
	"fmt"
	"net/http"
)

//func init()  {
//	router.AddRouter(router.RouteDetail{Name: "example",    Method:"GET",	Pattern:"/example",	HandlerFunc:HttpFuncExample})
//}


func HttpFuncExample(w http.ResponseWriter, req *http.Request)  {
	req.ParseForm()
	request_args := req.Form
	fmt.Printf(request_args.Encode())
	response := service.HttpServiceExample()
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(&response)
}
