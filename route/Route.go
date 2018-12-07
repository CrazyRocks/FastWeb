package route

import (
	"../app/controllers"
	"github.com/kataras/iris"
)

type Route struct {
	app iris.Application
}

func New() Route {
	return Route{}
}
func (route *Route) Init(app iris.Application) *Route {
	app.Get("/", new(controller.IndexController).Index)
	app.Get("/home", new(controller.IndexController).Home)
	route.app = app
	return route
}
