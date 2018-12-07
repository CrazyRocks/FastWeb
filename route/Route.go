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
	app.Get("/", before, controller.NewIndex().Index, after)
	app.Get("/home", controller.NewIndex().Home)
	route.app = app
	return route
}

func before(ctx iris.Context) {
	ctx.Writef("before action")
	ctx.Next()
}

func after(ctx iris.Context) {
	ctx.Writef("after action")
	ctx.Next()
}
