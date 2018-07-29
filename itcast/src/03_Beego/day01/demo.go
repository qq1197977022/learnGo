package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (mc *MainController) Get() {
	mc.Ctx.WriteString("HelloWorld")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
