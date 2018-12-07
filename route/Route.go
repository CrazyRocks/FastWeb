package route

import (
	"github.com/kataras/iris"
	"../app/controllers"
)

type Route struct {
	app iris.Application
}

func (route *Route) Init(app iris.Application) *Route {
	app.Get("/", new(controller.IndexController).Index)
	app.Get("/home", new(controller.IndexController).Home)
	route.app = app
	return route
}
