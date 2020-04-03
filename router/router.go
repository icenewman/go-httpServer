package router

import (
	"../handler"
	"github.com/gorilla/mux"
	"net/http"
	"../service"
)


type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//路由表，类似工厂模式用于注册路径，在此添加路由
var routes = Routes{
	Route{Name:"GetWeather",     Method:"GET",	Pattern:"/example",	HandlerFunc: handler.HttpFuncExample},
	Route{Name:"Health", Method:"GET",	Pattern:"/health",	HandlerFunc: handler.Health},
	Route{Name:"Check",  Method:"GET",   Pattern:"/check",	HandlerFunc: handler.ConsulCheck},
}

//批量注册路由
func Router() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = service.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
