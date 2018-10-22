package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"restapi"
)

type GatewayTableController struct {
	beego.Controller
}

type GWs struct {
	Result []restapi.Gateway
}

func (c *GatewayTableController) Get() {
	var gateways GWs

	gw := restapi.GetAllGateways("10")
	for i := 0; i < gw.TotalCount; i++ {
		gw.Result[i].Activity = restapi.GetGatewayActivity(gw.Result[i].MAC)
	}
	gateways.Result = gw.Result

	b, err := json.Marshal(gateways)
	if err != nil {
		log.Warn("gateway table json marshal error: %v", err.Error())
	}
	//fmt.Fprintf(c.Ctx.ResponseWriter, string(b))

	_, err = c.Ctx.ResponseWriter.Write(b)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (c *GatewayTableController) Post() {
	//c.TplName = "lora.html"

	mac := c.Input().Get("mac")

	gatewayDetails := restapi.GetGatewayDetails(mac)
	if gatewayDetails == nil {
		logs.Error("gatewayDetails is nil")
		c.Abort("500")
	}

	b, err := json.Marshal(gatewayDetails)
	if err != nil {
		log.Warn("gateway table json marshal error: %v", err.Error())
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))

}