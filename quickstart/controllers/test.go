package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.TplName = "new.html"
}

func (c *TestController) Post() {
	l := c.Ctx.Request.Form["linkType"]
	fmt.Println(l[0])

	link := c.Input().Get("linkType")
	fmt.Println("link: ", link)

	fmt.Fprint(c.Ctx.ResponseWriter, "success")
}
