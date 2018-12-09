package main

import (
	"./app"
	"github.com/kataras/iris"
)

func main() {
	irisApp := iris.New()
	irisApp.RegisterView(iris.HTML("./public/views", ".html"))
	app.GetIrisBoot().Boot(irisApp)
	irisApp.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/main.yml")))
}
