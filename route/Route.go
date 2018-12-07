package route

import (
	"../app/controllers"
	"github.com/kataras/iris"
)

type Route struct {
	app iris.Application
}

func New() *Route {
	return &Route{}
}

func (route *Route) Init(app iris.Application) *Route {
	app.Get("/", controller.NewIndex().Index)
	app.Get("/home", controller.NewIndex().Home)
	route.app = app
	return route
}
