package controllers

import (
	"github.com/astaxie/beego"
	"modifyfile"
	"netconfig"
	"strconv"
)

type Tab3Controller struct {
	beego.Controller
}

func (c *Tab3Controller) Post() {
	hostname := c.Input().Get("HostName")
	netconfig.ChangeHostName(hostname)

	localIP := c.Input().Get("localIP_1") + "." + c.Input().Get("localIP_2") + "." + c.Input().Get("localIP_3") + "." + c.Input().Get("localIP_4")
	localMask := "255.255." + c.Input().Get("localMask_1") + "." + c.Input().Get("localMask_2")
	modifyfile.ChangeFileParameter(WIRELESS_FN, "address", localIP)
	modifyfile.ChangeFileParameter(WIRELESS_FN, "netmask", localMask)
	//reboot() or networkRestart()

	startIP := "192.168." + c.Input().Get("startIPThird") + "." + c.Input().Get("startIPFourth")
	Range := c.Input().Get("maxRange")
	r, _ := strconv.Atoi(Range)
	netconfig.SetStartStopIP(startIP, r, c.Input().Get("startIPThird"))

	netconfig.SetRentTime(c.Input().Get("rentTime"))

	DNS1 := c.Input().Get("DHCPDNS1_1") + "." + c.Input().Get("DHCPDNS1_2") + "." + c.Input().Get("DHCPDNS1_3") + "." + c.Input().Get("DHCPDNS1_4")
	DNS2 := c.Input().Get("DHCPDNS2_1") + "." + c.Input().Get("DHCPDNS2_2") + "." + c.Input().Get("DHCPDNS2_3") + "." + c.Input().Get("DHCPDNS2_4")
	twoDNS := DNS1 + ", " + DNS2
	netconfig.SetDHCPDNS(twoDNS)

	netconfig.Reboot()
	//c.Redirect("/connection", 301)
}

