package router

import (
	"modifyfile"
	"os/exec"
	"fmt"
	"strings"
)

//write RIP configuration to profile(path: /etc/quagga/ripd.conf), only reboot will make it work
func StartRip() {
	configRipDaemons()
	configRipNetwork()
	QuaggaRestart()
}

func configRipDaemons() {
	FN := "/etc/quagga/daemons"
	modifyfile.ChangeFileParameter(FN, "ospfd", "no")
	modifyfile.ChangeFileParameter(FN, "ripd", "yes")
}

func configRipNetwork() {
	FN := "/etc/quagga/ripd.conf"
	net := IPtoNet(GetIP())
	modifyfile.ChangeFileParameter(FN, "network", net)
}

func GetIP()  []byte {
	var b []byte

	cmd := exec.Command("ifconfig", "eth0")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	l := strings.Index(string(stdout), "inet addr") + 10
	for i := l; stdout[i] != ' '; i++ {
		b = append(b, stdout[i])
	}

	return b
}

func IPtoNet(b []byte) string {
	bb := strings.Split(string(b), ".")
	net := bb[0] + "." + bb[1] + "." + bb[2] + "." + "0" + "/24"

	return net
}

func QuaggaRestart() {
	cmd := exec.Command("service", "quagga", "restart")
	cmd.Run()
}