package controllers

import (
	"github.com/astaxie/beego"
	"netconfig"
	"strings"
)

//const ETH_PORT = "eth0"

type ConnectionController struct {
	beego.Controller
}

func (c *ConnectionController) Get() {
	c.TplName = "connection.html"

	/*************************Tab-1***************************/
	wirelessName := netconfig.GetWirelessName()
	wirelessPasswd := netconfig.GetWirelessPasswd()
	c.Data["wirelessName"] = string(wirelessName)
	c.Data["wirelessPasswd"] = string(wirelessPasswd)

	/*************************Tab-2***************************/
	var ipConf netconfig.IPConfiguration
	ipConf.PortName = ETH_PORT
	ip, mask :=ipConf.GetIPandMask()
	c.Data["ip_1"] = ip[0]
	c.Data["ip_2"] = ip[1]
	c.Data["ip_3"] = ip[2]
	c.Data["ip_4"] = ip[3]
	c.Data["mask_1"] = mask[0]
	c.Data["mask_2"] = mask[1]
	c.Data["mask_3"] = mask[2]
	c.Data["mask_4"] = mask[3]

	g := netconfig.GetGateway()
	gw := strings.Split(string(g), ".")
	c.Data["gateway_1"] = gw[0]
	c.Data["gateway_2"] = gw[1]
	c.Data["gateway_3"] = gw[2]
	c.Data["gateway_4"] = gw[3]

	d := netconfig.GetDNS()
	dns := strings.Split(string(d), ".")
	c.Data["DNS1_1"] = dns[0]
	c.Data["DNS1_2"] = dns[1]
	c.Data["DNS1_3"] = dns[2]
	c.Data["DNS1_4"] = dns[3]
	c.Data["DNS2_1"] = dns[4]
	c.Data["DNS2_2"] = dns[5]
	c.Data["DNS2_3"] = dns[6]
	c.Data["DNS2_4"] = dns[7]

	domain := netconfig.GetDomainName()
	c.Data["domain"] = string(domain)

	c.Data["MTU"] = "1500"

	/*************************Tab-3***************************/
	hostname := netconfig.GetHostName()
	c.Data["hostname"] = string(hostname)

	wirelessIP, wirelessMask := netconfig.GetWirelessIPandMask()
	c.Data["wirelessIP"] = string(wirelessIP)
	c.Data["wirelessMask"] = string(wirelessMask)
	wIP := strings.Split(string(wirelessIP), ".")
	wMask := strings.Split(string(wirelessMask), ".")
	c.Data["localIP_1"] = wIP[0]
	c.Data["localIP_2"] = wIP[1]
	c.Data["localIP_3"] = wIP[2]
	c.Data["localIP_4"] = wIP[3]
	c.Data["localMask_1"] = wMask[2]
	c.Data["localMask_2"] = wMask[3]


	startIP, _ := netconfig.GetStartStopIP()
	s := strings.Split(string(startIP), ".")
	r := netconfig.GetRange()
	rentTime := netconfig.GetRentTime()
	c.Data["startIP_3"] = s[2]
	c.Data["startIP_4"] = s[3]
	c.Data["Range"] = r
	c.Data["rentTime"] = string(rentTime)

	DHCPDNS1, DHCPDNS2 := netconfig.GetDHCPDNS()
	dns1 := strings.Split(string(DHCPDNS1), ".")
	dns2 := strings.Split(string(DHCPDNS2), ".")
	c.Data["DHCPDNS1_1"] = dns1[0]
	c.Data["DHCPDNS1_2"] = dns1[1]
	c.Data["DHCPDNS1_3"] = dns1[2]
	c.Data["DHCPDNS1_4"] = dns1[3]
	c.Data["DHCPDNS2_1"] = dns2[0]
	c.Data["DHCPDNS2_2"] = dns2[1]
	c.Data["DHCPDNS2_3"] = dns2[2]
	c.Data["DHCPDNS2_4"] = dns2[3]
}