package controllers

import "github.com/astaxie/beego"

type WiFiController struct {
	beego.Controller
}

func (c *WiFiController) Get() {
	c.TplName = "wifi.html"
}
