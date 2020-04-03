package service

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
	"log"
)

type requestLog struct {
	RequestTime string `json:"request_time"`
	RequestAddr string `json:"request_addr"`
	RequestHost string `json:"request_host"`
	RequestArgs url.Values `json:"request_args"`
	RequestAgent string `json:"request_agent"`
}

//log记录
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inner.ServeHTTP(w, r)
		r.ParseForm()
		sessionLog := new(requestLog)
		now := time.Now().Format("2006-01-02 15:04:05")
		sessionLog.RequestTime = now
		sessionLog.RequestAddr = r.RemoteAddr
		sessionLog.RequestHost = r.Host
		sessionLog.RequestArgs = r.Form
		sessionLog.RequestAgent = r.UserAgent()

		data,err := json.Marshal(sessionLog)
		if err != nil{
			log.Println("log json Marshal fail")
		}
		log.Println(string(data))
	})
}
