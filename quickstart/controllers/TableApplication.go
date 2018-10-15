package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gpmgo/gopm/modules/log"
	"restapi"
)

type ApplicationTableController struct {
	beego.Controller
}

type Apps struct {
	Result []restapi.Application
}

func (c *ApplicationTableController) Get() {
	var apps Apps
	appsList := restapi.GetAllApplications("10")
	apps.Result = appsList.Result
	b, err := json.Marshal(apps)
	if err != nil {
		log.Warn("application table json marshal error: %v", err.Error())
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}
