package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
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
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))


	/*b, err := ioutil.ReadFile("/home/h/goPath/src/quickstart/controllers/gwtable")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))*/
}

func (c *GatewayTableController) Post() {
	mac := c.Input().Get("mac")
	gatewayDetails := restapi.GetGatewayDetails(mac)

	c.TplName = "lora.html"

	fmt.Println(gatewayDetails.Altitude)

	c.Data["LoRaAltitude"] = "100"

}