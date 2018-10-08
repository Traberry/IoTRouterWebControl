package netconfig

import (
	"io/ioutil"
	"fmt"
	"os/exec"
	"os"
)

const HOSTNAME_FN = "/etc/hostname"

func DeleteFile(path string) {
	_, err := os.Stat(path)
	if err == nil {
		os.Remove(path)
	}
}

func UnzipFile(fromFile, toFile string) {
	cmd := exec.Command("unzip", "-o", fromFile, "-d", toFile)
	cmd.Run()
}

func GetHostName() []byte {
	cmd := exec.Command("cat", "/etc/hostname")
	//cmd := exec.Command("hostname")
	b, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func ChangeHostName(hostname string) {
	err := ioutil.WriteFile(HOSTNAME_FN, []byte(hostname), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func SetTimeZone_1() {
	cmd1 := exec.Command("rm", "/etc/localtime")
	cmd1.Run()

	//England
	cmd2 := exec.Command("ln", "-s", "/usr/share/zoneinfo/GMT0", "/etc/localtime")
	cmd2.Run()
}

func SetTimeZone_12() {
	cmd1 := exec.Command("rm", "/etc/localtime")
	cmd1.Run()

	//Shanghai, China
	cmd2 := exec.Command("ln", "-s", "/usr/share/zoneinfo/Asia/Shanghai", "/etc/localtime")
	cmd2.Run()
}




func NetworkRestart() {
	cmd := exec.Command("service", "networking", "restart")
	cmd.Run()
}

func Reboot() {
	cmd := exec.Command("reboot")
	cmd.Run()
}