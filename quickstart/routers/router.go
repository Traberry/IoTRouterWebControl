package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/login", &controllers.LogController{})
    beego.Router("/", &controllers.MainController{})
	beego.Router("/connection", &controllers.ConnectionController{})
	beego.Router("/connection/tab1", &controllers.Tab1Controller{})
    beego.Router("/connection/tab2", &controllers.Tab2Controller{})
	beego.Router("/connection/tab3", &controllers.Tab3Controller{})
	beego.Router("/connection/tab4", &controllers.Tab4Controller{})
	beego.Router("/wifi", &controllers.WiFiController{})
	beego.Router("/lora", &controllers.LoRaController{})
	beego.Router("/lora/organizationTable", &controllers.OrgTableController{})
	beego.Router("/lora/userTable", &controllers.UserTableController{})
	beego.Router("/lora/gatewayTable", &controllers.GatewayTableController{})
	beego.Router("/lora/serverTable", &controllers.ServerTableController{})
	beego.Router("/lora/applicationTable", &controllers.ApplicationTableController{})
	beego.Router("/lora/applicationTable/deviceOfAppID/?:id", &controllers.DeviceOfAppTableController{})
	beego.Router("/lora/deviceTable", &controllers.DeviceTableController{})

}
