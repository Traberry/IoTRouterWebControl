package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gpmgo/gopm/modules/log"
	"restapi"
)

type DeviceActivationController struct {
	beego.Controller
}

func (c * DeviceActivationController) Post() {
	devEUI := c.Input().Get("devEUI")
	devActive := restapi.GetDeviceActivation(devEUI)
	if devActive == nil {
		logs.Error("gatewayDetails is nil")
		c.Abort("500")
	}

	b, err := json.Marshal(devActive)
	if err != nil {
		log.Warn("device activation json marshal error: %v", err.Error())
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}
