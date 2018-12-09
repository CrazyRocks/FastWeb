package core

import (
	"github.com/kataras/iris"
)

/**
控制器中间件
*/
type ControllerMiddleware struct {
	Name        string
	Description string
	Middlewares map[string]IMiddleware
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

func (mid *ControllerMiddleware) RegisterMiddlewares(midNames []string) {
}
