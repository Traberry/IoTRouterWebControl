package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"restapi"
	"strconv"
)

type DeviceTableController_1 struct {
	beego.Controller
}

func (c *DeviceTableController_1) Get() {
	type app struct {
		Name string `json:"name"`
		Number string `json:"number"`
	}
	type list struct {
		Result []app
	}

	var li list
	appList := restapi.GetAllApplications("10")
	count, _ := strconv.Atoi(appList.TotalCount)
	for i := 0; i < count; i++ {
		li.Result = append(li.Result, app{Name:(*appList).Result[i].Name, Number: (*appList).Result[i].ID})
	}

	b, err := json.Marshal(li.Result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}

