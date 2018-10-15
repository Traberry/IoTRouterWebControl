package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gpmgo/gopm/modules/log"
	"restapi"
)

type ServerTableController struct {
	beego.Controller
}

func (c *ServerTableController) Get() {
	s := restapi.GetServers("10")
	b, err := json.Marshal(s)
	if err != nil {
		log.Warn("server table json marshal error: %v", err.Error())
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}