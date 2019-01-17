package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/login", &controllers.LogController{})
    beego.Router("/", &controllers.MainController{})

    beego.Router("/jump", &controllers.JumpController{})

    beego.Router("/test", &controllers.TestController{})

	beego.Router("/connection", &controllers.ConnectionController{})
	beego.Router("/connection/tab1", &controllers.Tab1Controller{})
    beego.Router("/connection/tab2", &controllers.Tab2Controller{})
	beego.Router("/connection/tab3", &controllers.Tab3Controller{})
	beego.Router("/connection/tab4", &controllers.Tab4Controller{})
	beego.Router("/connection/staticRouteTable", &controllers.RouterTableController{})
	beego.Router("/wifi", &controllers.WiFiController{})

	beego.Router("/lora", &controllers.LoRaController{})
	beego.Router("/lora/organizationTable", &controllers.OrgTableController{})
	beego.Router("/lora/gatewayTable", &controllers.GatewayTableController{})
	beego.Router("/lora/serverTable", &controllers.ServerTableController{})
	beego.Router("/lora/applicationTable", &controllers.ApplicationTableController{})
	beego.Router("/lora/applicationTable/deviceOfAppID/?:id", &controllers.DeviceOfAppTableController{})
	beego.Router("/lora/deviceTableApps", &controllers.DeviceTableController_1{})
	beego.Router("/lora/deviceTableDevices/?:id", &controllers.DeviceTableController_2{})
	beego.Router("/lora/deviceActivation", &controllers.DeviceActivationController{})
	beego.Router("/lora/topology", &controllers.TopoController{})

	beego.Router("/user", &controllers.UserManagementController{})
	beego.Router("/user/judgeUser", &controllers.UserJudgeController{})
	beego.Router("/user/userList", &controllers.UserTableController{})
	beego.Router("/user/organizationIDs", &controllers.OrgIDController{})
	beego.Router("/user/addUser", &controllers.UserAddController{})
	beego.Router("/user/modifyUser", &controllers.UserUpdateController{})
	beego.Router("/user/deleteUser", &controllers.UserDeleteController{})

	beego.Router("/lora/gatewayLog", &controllers.GatewayLogController{})

}
