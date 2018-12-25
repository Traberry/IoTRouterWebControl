package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"restapi"
)

type JumpController struct {
	beego.Controller
}

func (c *JumpController) Post() {
	ip := restapi.IPAddressOfAPIServer
	url := "https://" + ip
	status := restapi.LoRaAppServerLogIn(UserName, Password)
	if status == 200 {
		c.Redirect(url, 301)
	}else {
		fmt.Fprint(c.Ctx.ResponseWriter, string(status))
		log.Println("unauthorized")
	}
}