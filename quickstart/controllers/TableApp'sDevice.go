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

var id string

func (c *DeviceOfAppTableController) Get() {
	devices := restapi.GetSimpleDevices(id, "10")
	b, err := json.Marshal(devices)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id: ", id)
	fmt.Println(string(b))
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}

/*
 *the front table only show three columns: Name, LinkMargin, Battery
 */
func (c *DeviceOfAppTableController) Post() {
	id = c.Input().Get("id")//the front should give the id to the backend ???
	c.Ctx.WriteString(id)
}
