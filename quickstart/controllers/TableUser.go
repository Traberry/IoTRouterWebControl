package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gpmgo/gopm/modules/log"
	"restapi"
)

type UserTableController struct {
	beego.Controller
}

func (c *UserTableController) Get() {
	u := restapi.GetUsers("10")
	b, err := json.Marshal(u)
	if err != nil {
		log.Warn("user table json marshal error: %v", err.Error())
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}
