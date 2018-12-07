package core

import (
	"github.com/kataras/iris"
)

type BaseController struct {
	ModuleName  string
	Middlewares []interface{}
}

func (ctrl *BaseController) Index(ctx iris.Context) {
	ctx.JSON([]string{"test", "test2"})
}
