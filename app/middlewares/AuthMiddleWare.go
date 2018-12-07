package middlewares

import "github.com/kataras/iris"

func before(ctx iris.Context){
	println("before AuthMiddleWare")
	ctx.Next()
}

func after(ctx iris.Context){
	println("after AuthMiddleWare")
}
