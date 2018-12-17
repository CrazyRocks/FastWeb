package controller

import (
	"github.com/kataras/iris"
)

type IndexController struct {
	BaseController
}

func (ctrl *IndexController) Home(ctx iris.Context) {
	ctx.View("home.html")
}
