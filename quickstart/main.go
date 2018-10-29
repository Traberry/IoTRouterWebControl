package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.EnableHTTP = false

	beego.Run()
}