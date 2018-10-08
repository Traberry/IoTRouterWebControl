package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"restapi"
)

type UserTableController struct {
	beego.Controller
}

func (c *UserTableController) Get() {
	u := restapi.GetUsers("10")
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}
