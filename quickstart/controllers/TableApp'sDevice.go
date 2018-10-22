package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"restapi"
)

type DeviceOfAppTableController struct {
	beego.Controller
}

func (c *DeviceOfAppTableController) Get() {
	//fmt.Println("deviceOfApp request: ", c.Ctx.Request)

	id := c.Ctx.Input.Param(":id")
	devices := restapi.GetSimpleDevices(id, "10")
	b, err := json.Marshal(devices)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}
