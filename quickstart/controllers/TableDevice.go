package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"restapi"
	"strconv"
)

type DeviceTableController struct {
	beego.Controller
}

/*
 *show all applications and their devices
 */
func (c *DeviceTableController) Get() {
	//get all AppIDs
	var appID []string
	appList := restapi.GetAllApplications("10")
	count, _ := strconv.Atoi(appList.TotalCount)
	for i := 0; i < count; i++ {
		appID = append(appID, appList.Result[i].ID)
	}

	//return device list according to AppID every time
	for _, v := range appID {
		fmt.Printf("*****Application %s*****\n", v)
		deviceList := restapi.GetDevices( v, "10")
		b, err := json.Marshal(deviceList)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
	}
}