package router

import (
	"os/exec"
	"github.com/google/seesaw/quagga"
	"fmt"
)

/****add static route by using command line "route"****/
func AddStaticRouteToHost(hostIP string, gatewayIP string) {
	cmd := exec.Command("route", "add", "-host", hostIP, "gw", gatewayIP)
	cmd.Run()
}

func AddStaticRouteToNet(netAddr string, netMask string, gatewayIP string) {
	cmd := exec.Command("route", "add", "-net", netAddr, "netmask", netMask, "gw", gatewayIP)
	cmd.Run()
}

//write route configuration to zebra.conf
func writeToZebra() {
	//it is hard to write words directly to the configuration file "zebra.conf"
	//there is no parameter as the form key-value
}

/****adding static route by using quagga zebra, only reboot will make it work****/
func addStaticRoutebyQuagga() {
	path := "/var/run/quagga/zebra.vty"
	vty := quagga.NewVTY(path)
	err := vty.Dial()
	if err != nil {
		fmt.Println(err)
	}

	cmd3 := []string{"en", "conf t", "ip route 192.168.10.0/24 192.168.10.1", "write"}
	err2 := vty.Commands(cmd3)
	if err2 != nil {
		fmt.Println(err)
	}
}
