package main

import (
	"./route"
	"github.com/kataras/iris"
)

func main(){
	app := iris.New()
	new(route.Route).Init(*app)
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/main.yml")))
}
