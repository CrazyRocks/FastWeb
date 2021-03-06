package route

import (
	"../app/controllers"
	"github.com/kataras/iris"
	"sync"
)

type Route struct {
	app         iris.Application
}

var router *Route
var lock *sync.Mutex = &sync.Mutex{}

func New() *Route {
	return &Route{}
}

/**
获取路由实例
*/
func GetInstance() *Route {
	if router == nil {
		lock.Lock()
		defer lock.Unlock()
		if router == nil {
			router = &Route{}
		}
	}
	return router
}

/**
初始化路由
*/
func (route *Route) Boot(app iris.Application) *Route {
	indexController := controller.NewIndex()
	indexController.RegisterMiddleware("auth")
	app.Get("/", indexController.Before, indexController.Index, indexController.After)
	app.Get("/home", indexController.Home)

	userController := controller.NewUser()
	userController.RegisterMiddleware("auth", "db")
	app.Post("/users", userController.Before, userController.Add, userController.After)

	route.app = app
	return route
}
