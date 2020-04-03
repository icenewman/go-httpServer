package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthData struct {
	Statu string
	Msg   string
}

func Health(w http.ResponseWriter, req *http.Request) {
	heath := HealthData{"UP","OK"}
	json.NewEncoder(w).Encode(&heath)
}


func ConsulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}