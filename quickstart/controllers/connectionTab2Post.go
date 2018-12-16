package controllers

import (
	"github.com/astaxie/beego"
	"netconfig"
)

type Tab2Controller struct {
	beego.Controller
}

func (c *Tab2Controller) Post() {
	linkType := c.Input().Get("linkType")

	var ipConf netconfig.IPConfiguration
	if linkType == "0" {
		ipConf.IP = c.Input().Get("ip_1") + "." + c.Input().Get("ip_2") + "." + c.Input().Get("ip_3") + "." + c.Input().Get("ip_4")
		ipConf.Mask = c.Input().Get("mask_1") + "." + c.Input().Get("mask_2") + "." + c.Input().Get("mask_3") + "." + c.Input().Get("mask_4")
		ipConf.GW = c.Input().Get("gateway_1") + "." + c.Input().Get("gateway_2") + "." + c.Input().Get("gateway_3") + "." + c.Input().Get("gateway_4")
		ipConf.PrimaryDNS = c.Input().Get("DNS1_1") + "." + c.Input().Get("DNS1_2") + "." + c.Input().Get("DNS1_3") + "." + c.Input().Get("DNS1_4")
		ipConf.OptionalDNS = c.Input().Get("DNS2_1") + "." + c.Input().Get("DNS2_2") + "." + c.Input().Get("DNS2_3") + "." + c.Input().Get("DNS2_4")
		ipConf.MTU = c.Input().Get("MTU")
		ipConf.PortName = ETH_PORT

		ipConf.SetIP()
		ipConf.SetGW()
		ipConf.SetMTU()
		ipConf.SetPrimaryDNS()
		ipConf.SetOptionalDNS()
	}

	if linkType == "1" {
		ipConf.MTU = c.Input().Get("MTU")
		ipConf.PortName = ETH_PORT
	}

	netconfig.Reboot()
	//c.Redirect("/connection", 301)
}
