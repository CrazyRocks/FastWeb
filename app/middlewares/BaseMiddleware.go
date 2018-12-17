package middlewares

import "github.com/kataras/iris"

//基础中间件
type BaseMiddleware struct {
	Name        string
	Description string
}

func (mid *BaseMiddleware) Before(ctx iris.Context) {

}

func (mid *BaseMiddleware) After(ctx iris.Context) {

}
