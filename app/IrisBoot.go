package app

import (
	"../route"
	"./middlewares"
	"github.com/kataras/iris"
	"sync"
)

type IrisBoot struct {
	Middlewares map[string]interface{}
	App         iris.Application
	isBooted    bool
}

var irisBootInstance *IrisBoot
var lock *sync.Mutex = &sync.Mutex{}

func GetIrisBoot() *IrisBoot {
	if irisBootInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if irisBootInstance == nil {
			irisBootInstance = &IrisBoot{}
		}
	}
	return irisBootInstance
}

func (irisBoot *IrisBoot) Boot(app *iris.Application) {
	if irisBoot.isBooted {
		return
	}
	irisBoot.App = *app
	irisBoot.registerMiddlewares()
	irisBoot.registerRoutes()
}

func (irisBoot *IrisBoot) registerMiddlewares() {
	irisBoot.Middlewares = make(map[string]interface{})
	irisBoot.Middlewares["auth"] = &middlewares.AuthMiddleware{}
}

func (irisBoot *IrisBoot) registerRoutes() {
	router := route.GetInstance()
	router.Boot(irisBoot.App)
}
