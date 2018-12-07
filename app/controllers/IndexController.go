package controller

import (
	"../core"
	"github.com/kataras/iris"
)

type IndexController struct {
	*core.BaseController
}

func (ctrl *IndexController) Home(ctx iris.Context) {
	ctx.Writef("This is Home Page")
}
