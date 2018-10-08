package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"restapi"
)

type ServerTableController struct {
	beego.Controller
}

func (c *ServerTableController) Get() {
	s := restapi.GetServers("10")
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}