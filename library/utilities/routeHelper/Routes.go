package routeHelper

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name         string
	Method       string
	Pattern      string
	Description  string
	AuthRequired bool
	URLParams    interface{}
	BodyParams   interface{}
	HandlerFunc  http.HandlerFunc
}

func (self *Route) Add(Name, Method, Pattern, Description string, AuthRequired bool, URLParams, BodyParams interface{}) {
	self.Name = Name
	self.Method = Method
	self.Pattern = Pattern
	self.Description = Description
	self.AuthRequired = AuthRequired
	self.URLParams = URLParams
	self.BodyParams = BodyParams
}

func (self *Route) Output() (data RouteDescription) {
	data.Name = self.Name
	data.Method = self.Method
	data.Pattern = self.Pattern
	data.Description = self.Description
	data.AuthRequired = self.AuthRequired
	data.URLParams = self.URLParams
	data.BodyParams = self.BodyParams
	return
}

type Routes struct {
	list []Route
}

func (self *Routes) Add(Name, Method, Pattern, Description string, AuthRequired bool, URLParams, BodyParams interface{}, HandlerFunc http.HandlerFunc) {
	self.list = append(self.list, Route{
		Name,
		Method,
		Pattern,
		Description,
		AuthRequired,
		URLParams,
		BodyParams,
		HandlerFunc,
	})
}

//add multiple routes that uses the same handler function to list
func (self *Routes) AddObject(HandlerFunc http.HandlerFunc, routes ... Route) {
	for i := 0; i < len(routes); i++ {
		routes[i].HandlerFunc = HandlerFunc
		self.list = append(self.list, routes[i])
	}
}

func NewRouter(data Routes) *mux.Router {
	var helpHandler Handler
	helpHandler.Init(data.list)
	helper := Route{
		"Help",
		"GET",
		"/help",
		"List api description on this microservice",
		false,
		nil,
		nil,
		helpHandler.Help,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods(helper.Method).
		Path(helper.Pattern).
		Name(helper.Name).
		Handler(helper.HandlerFunc)
	for _, route := range data.list {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
