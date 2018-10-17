package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
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


	/*b, err := ioutil.ReadFile("/home/h/goPath/src/quickstart/controllers/gwtable")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))*/
}

func (c *GatewayTableController) Post() {
	fmt.Println("method: ", c.Ctx.Request.Method)
	fmt.Println("url: ", c.Ctx.Request.URL)
	fmt.Println("all: ", c.Ctx.Request)
	mac := c.GetString("mac")//how to get mac
	fmt.Println("mac: ", mac)
	gatewayDetails := restapi.GetGatewayDetails(mac)
	fmt.Println(gatewayDetails)

	//c.TplName = "lora.html"

	fmt.Println("gaodu: ", gatewayDetails.Altitude)

	c.Data["LoRaAltitude"] = 100

}