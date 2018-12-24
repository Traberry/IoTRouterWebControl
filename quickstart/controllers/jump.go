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
	status := restapi.LoRaAppServerLogIn("admin", "admin")
	if status == 200 {
		c.Redirect("https://192.168.3.105:8080", 301)
	}else {
		fmt.Fprint(c.Ctx.ResponseWriter, string(status))
		log.Println("unauthorized")
	}
}