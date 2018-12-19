package controllers

import "github.com/astaxie/beego"

type LoRaController struct {
	beego.Controller
}

func (c *LoRaController) Get() {
	//judge user login first
	if !LogIn {
		c.Redirect("/login", 301)
	}

	c.TplName = "lora.html"
}
