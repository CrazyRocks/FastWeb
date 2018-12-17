package controller

import (
	"../core"
	"../middlewares"
	"github.com/kataras/iris"
)

/**
控制器中间件
*/
type ControllerMiddleware struct {
	Name        string
	Description string
	Middlewares map[string]core.IMiddleware
}

func (mid *ControllerMiddleware) Before(ctx iris.Context) {
	if mid.Middlewares != nil {
		for _, handler := range mid.Middlewares {
			ctx.AddHandler(handler.Before)
		}
	}
	ctx.Next()
}

func (mid *ControllerMiddleware) After(ctx iris.Context) {
	if mid.Middlewares != nil {
		for _, handler := range mid.Middlewares {
			ctx.AddHandler(handler.After)
		}
	}
	ctx.Next()
}

func (mid *ControllerMiddleware) RegisterMiddleware(midNames ...string) {
	if len(midNames) == 0 {
		return
	}
	if mid.Middlewares == nil {
		mid.Middlewares = make(map[string]core.IMiddleware)
	}
	for _, midName := range midNames {
		mid.Middlewares[midName] = middlewares.RegisterMiddlewares()[midName]
	}
}
