package core

import (
	"github.com/kataras/iris"
)

type BaseController struct {
	ModuleName string
}

func (ctrl *BaseController) Index(ctx iris.Context) {
	ctx.JSON([]string{"test", "test2"})
	//ctx.Writef("This is Base.index")
}
