package app

import (
	"../route"
	"github.com/kataras/iris"
	"sync"
)

type IrisBoot struct {
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
	irisBoot.registerRoutes()
}

func (irisBoot *IrisBoot) registerRoutes() {
	router := route.GetInstance()
	router.Boot(irisBoot.App)
}
