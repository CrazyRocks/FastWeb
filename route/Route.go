package route

import (
	"../app/controllers"
	"github.com/kataras/iris"
	"sync"
)

type Route struct {
	app         iris.Application
	Middlewares map[string]interface{}
	Events      map[string]interface{}
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
	indexController.RegisterMiddlewares([]string{"auth"})
	app.Get("/", indexController.Before, indexController.Index, indexController.After)
	app.Get("/home", indexController.Home)
	route.app = app
	return route
}
