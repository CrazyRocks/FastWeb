package middlewares

import (
	"github.com/kataras/iris"
)

type AuthMiddleware struct {
	BaseMiddleware
}

func (mid *AuthMiddleware) Before(ctx iris.Context) {
	mid.BaseMiddleware.Before(ctx)
	println("before AuthMiddleWare")
	ctx.Next()
}

func (mid *AuthMiddleware) After(ctx iris.Context) {
	mid.BaseMiddleware.After(ctx)
	println("after AuthMiddleWare")
	ctx.Next()
}
