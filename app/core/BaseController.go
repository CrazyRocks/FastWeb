package core

import (
	"github.com/kataras/iris"
)

type BaseController struct {
	ModuleName string
	ControllerMiddleware
}

func (ctrl *BaseController) Index(ctx iris.Context) {
	ctx.JSON([]string{"wendao", "17173"})
	ctx.Next()
}
