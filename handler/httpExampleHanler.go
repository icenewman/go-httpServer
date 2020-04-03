package handler

import (
	"../service"
	"encoding/json"
	"net/http"
	"net/url"
)


func HttpFuncExample(w http.ResponseWriter, req *http.Request)  {
	req.ParseForm()
	request_args := req.Form
	response := service.httpServiceExample()
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(&response)
}
