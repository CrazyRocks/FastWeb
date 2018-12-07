package core

import (
	"encoding/json"
	"github.com/kataras/iris"
)

type BaseController struct {
	ModuleName string
}

func (ctrl *BaseController) New() {
}

func (ctrl *BaseController) Index(ctx iris.Context){
	ctrl.Return([]string{"test", "test2"}, "json", ctx)
	//ctx.Writef("This is Base.index")
}

func (ctrl *BaseController) Return(data interface{}, returnType string, ctx iris.Context){
	if(returnType == "json"){
		returnData, _ := json.Marshal(data)
		ctx.Header("Content-Type", "application/json")
		ctx.Writef(string(returnData))
	}

}
