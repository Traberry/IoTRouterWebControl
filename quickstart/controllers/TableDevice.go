package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"restapi"
)

type DeviceTableController_2 struct {
	beego.Controller
}
/*
 *
 */
func (c *DeviceTableController_2) Get() {
	id := c.Ctx.Input.Param(":id")
	deviceList := restapi.GetDevices(id, "100")
	b, err := json.Marshal(deviceList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}