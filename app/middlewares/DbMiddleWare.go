package middlewares

import (
	"github.com/kataras/iris"
	"../models/db"
)

type DbMiddleware struct {
	BaseMiddleware
}

func (mid *DbMiddleware) Before(ctx iris.Context) {
	mid.BaseMiddleware.Before(ctx)
	println("before DbMiddleware")
	ctx.Next()
}

func (mid *DbMiddleware) After(ctx iris.Context) {
	mid.BaseMiddleware.After(ctx)
	println("after DbMiddleware")
	db.Db().Close()
	ctx.Next()
}
