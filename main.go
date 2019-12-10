package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"smyappTwo/route"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	route.Route(app)
	app.Run(iris.Addr(":9999"))

}


func Role(app *mvc.Application){

}