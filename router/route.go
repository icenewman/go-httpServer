package router

import (
	"../tool/log"
	"github.com/gorilla/mux"
	"net/http"
)


type RouteList []RouteDetail

type RouteDetail struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}


type RouteMapInterface interface {
}


func (r *RouteList)AddHttpRoute(name string, method string, routePath string, serviceName http.HandlerFunc)  {
	item := RouteDetail{Name:name, Method:method, Pattern:routePath, HandlerFunc:serviceName}
	*r = append(*r, item)
}

//批量注册路由
func RouterFilter(routes RouteList) *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = log.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func AddRouter(eface RouteDetail){
	Routes = append( Routes, eface)
	MuxRouter = RouterFilter(Routes)
}
