package middlewares

import (
	"../core"
)

var middlewareMap map[string] core.IMiddleware
var isBooted bool

func Boot(){
	middlewareMap = make(map[string]core.IMiddleware)
	middlewareMap["auth"] = &AuthMiddleware{}
	isBooted = true
}

func RegisterMiddlewares() map[string] core.IMiddleware {
	if !isBooted {
		Boot()
	}
	return middlewareMap
}
