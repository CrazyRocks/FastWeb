package core

import "github.com/kataras/iris"

type IMiddleware interface {
	Before(ctx iris.Context)
	After(ctx iris.Context)
}
