package router

import (
	"../handler"
	"github.com/gorilla/mux"
)


var Routes RouteList
var MuxRouter *mux.Router

//路由注册
func init(){
	AddRouter(RouteDetail{Name: "Health",    Method:"GET",	Pattern:"/health",	HandlerFunc:handler.Health})
	AddRouter(RouteDetail{Name: "example",    Method:"GET",	Pattern:"/example",	HandlerFunc:handler.HttpFuncExample})
}