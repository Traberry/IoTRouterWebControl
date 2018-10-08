package netconfig

import (
	"os/exec"
	"fmt"
	"strings"
	"modifyfile"
)

const WIRELESS_AP_FN = "/etc/hostapd/hostapd.conf"

//get wireless port's IP and mask information, it returns two byte slices
func GetWirelessIPandMask() ([]byte, []byte) {
	var b []byte
	var m []byte

	cmd := exec.Command("ifconfig", "wlan0")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	l1 := strings.Index(string(stdout), "inet addr") + 10
	for i := l1; stdout[i] != ' '; i++ {
		b = append(b, stdout[i])
	}

	l2 := strings.Index(string(stdout), "Mask") + 5
	for i := l2; stdout[i] != '\n'; i++ {
		m = append(m, stdout[i])
	}

	return b, m
}

func GetWirelessName() []byte {
	return modifyfile.ReadFileParameter(WIRELESS_AP_FN, "ssid")
}

func GetWirelessPasswd() []byte {
	return modifyfile.ReadFileParameter(WIRELESS_AP_FN, "wpa_passphrase")
}
