package controllers

import "github.com/astaxie/beego"

type LoRaController struct {
	beego.Controller
}

func (c *LoRaController) Get() {
	c.TplName = "lora.html"
}
