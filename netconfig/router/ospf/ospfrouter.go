package router

import (
	"modifyfile"
	"netconfig/router/rip"
)

func StartOSPF(area string) {
	configOSPFDaemons()
	configOSPFNetwork(area)
	router.QuaggaRestart()
}

func configOSPFDaemons() {
	FN := "/etc/quagga/daemons"
	modifyfile.ChangeFileParameter(FN, "ospfd", "yes")
	modifyfile.ChangeFileParameter(FN, "ripd", "no")
}

func configOSPFNetwork(a string) {
	FN := "/etc/quagga/ospfd.conf"
	IP := router.GetIP()
	net := router.IPtoNet(IP)
	modifyfile.ChangeFileParameter(FN, "ospf router-id", string(IP) )
	modifyfile.ChangeFileParameter(FN, "network", net)
	modifyfile.ChangeFileParameter(FN, "area", a)
}
