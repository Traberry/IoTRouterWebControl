package controllers

import (
	"github.com/astaxie/beego"
	"netconfig"
)

const (
	ETH_PORT = "eth0"
	WIRELESS_AP_FN = "/etc/hostapd/hostapd.conf"
	WIRELESS_FN = "/etc/network/interfaces"
	HOSTNAME_FN = "/etc/hostname"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"

	var eth netconfig.NetTransStatistics
	eth.GetEthSta()
	c.Data["ethRxPackets"] = eth.RxPackets
	c.Data["ethRxBytes"] = eth.RxBytes
	c.Data["ethRxErrors"] = eth.RxErrors
	c.Data["ethRxDropped"] = eth.RxDropped
	c.Data["ethTxPackets"] = eth.TxPackets
	c.Data["ethTxBytes"] = eth.TxBytes
	c.Data["ethTxErrors"] = eth.TxErrors
	c.Data["ethTxDropped"] = eth.TxDropped
	c.Data["ethTxCollisions"] = eth.TxCollisions

	var wlan netconfig.NetTransStatistics
	wlan.GetWlanSta()
	c.Data["wlanRxPackets"] = wlan.RxPackets
	c.Data["wlanRxBytes"] = wlan.RxBytes
	c.Data["wlanRxErrors"] = wlan.RxErrors
	c.Data["wlanRxDropped"] = wlan.RxDropped
	c.Data["wlanTxPackets"] = wlan.TxPackets
	c.Data["wlanTxBytes"] = wlan.TxBytes
	c.Data["wlanTxErrors"] = wlan.TxErrors
	c.Data["wlanTxDropped"] = wlan.TxDropped
	c.Data["wlanTxCollisions"] = wlan.TxCollisions
}