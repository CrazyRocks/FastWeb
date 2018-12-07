package main

import (
	"./route"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./public/views", ".html"))
	new(route.Route).Init(*app)
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/main.yml")))
}
