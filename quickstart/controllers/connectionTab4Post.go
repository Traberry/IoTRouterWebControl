package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"netconfig"
	"netconfig/router/static"
	router2 "netconfig/router/ospf"
	router3 "netconfig/router/rip"
)

type Tab4Controller struct {
	beego.Controller
}

func (c *Tab4Controller) Get() {
	//show static route table


	//fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}

func (c *Tab4Controller) Post() {
	routeID := c.Input().Get("NAT")
	switch routeID {
	case "0" :
		fmt.Println("NAT has already started")
	case "1":
		router3.StartRip()
	case "2":
		area := c.Input().Get("OSPFarea")
		router2.StartOSPF(area)
	default:
		fmt.Println("this message should not print in theory")
	}

	//handle static route to a net
	targetIP := c.Input().Get("targetIP_1") + "." + c.Input().Get("targetIP_2") + "." + c.Input().Get("targetIP_3") + "." + c.Input().Get("targetIP_4")
	targetMask := c.Input().Get("targetMask_1") + "." + c.Input().Get("targetMask_2") + "." + c.Input().Get("targetMask_3") + "." + c.Input().Get("targetMask_4")
	targetGateway := c.Input().Get("targetGateway_1") + "." + c.Input().Get("targetGateway_2") + "." + c.Input().Get("targetGateway_3") + "." + c.Input().Get("targetGateway_4")
	//router.AddStaticRouteToHost(targetIP, targetGateway)
	router.AddStaticRouteToNet(targetIP, targetMask, targetGateway)

	netconfig.Reboot()
	//c.Redirect("/connection", 301)
}
