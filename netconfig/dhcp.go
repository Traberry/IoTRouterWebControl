package netconfig

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"modifyfile"
)

const (
	DHCP_CONF = "/etc/dhcp/dhcpd.conf"
)

//get the start and the stop address of DHCP by reading its configuration file
//it returns two byte slices
func GetStartStopIP() ([]byte, []byte) {
	var startIP []byte
	var stopIP []byte

	b, err := ioutil.ReadFile(DHCP_CONF)
	if err != nil {
		fmt.Println(err)
	}

	l1 := strings.Index(string(b), "range") + 6
	for i := l1; b[i] != ' '; i++ {
		startIP = append(startIP, b[i])
	}

	l2 := l1 + len(startIP) +1
	for j := l2; b[j] != ';'; j++ {
		stopIP = append(stopIP, b[j])
	}

	return startIP, stopIP
}

//the stop IP's fourth number subtract the start IP's fourth number
func GetRange() int {
	start, stop := GetStartStopIP()
	s := strings.Split(string(start), ".")[3]
	t := strings.Split(string(stop), ".")[3]

	ss, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	tt, err := strconv.Atoi(t)
	if err != nil {
		fmt.Println(err)
	}

	return tt - ss
}

func GetRentTime() []byte {
	var time []byte

	b, err := ioutil.ReadFile(DHCP_CONF)
	if err != nil {
		fmt.Println(err)
	}
	l := strings.Index(string(b), "default-lease-time") + 19
	for i := l; b[i] != ';'; i++ {
		time = append(time, b[i])
	}

	return time
}

func GetDHCPDNS() ([]byte, []byte) {
	var DNS1 []byte
	var DNS2 []byte

	b, err := ioutil.ReadFile(DHCP_CONF)
	if err != nil {
		fmt.Println(err)
	}

	l1 := strings.Index(string(b), "option domain-name-servers") + 27
	for i := l1; b[i] != ','; i++ {
		DNS1 = append(DNS1, b[i])
	}

	l2 := l1 + len(DNS1) + 2
	for j := l2; b[j] != ';'; j++ {
		DNS2 = append(DNS2, b[j])
	}

	return DNS1, DNS2
}

//set start and stop IP address of DHCP server by change its configuration file
//the second parameter "r" means range
//the third parameter "m" means ip_3
func SetStartStopIP(startIP string, r int, m string) {
	s := strings.Split(startIP, ".")[3]
	ss, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	tt := strconv.Itoa(ss+r)
	stopIP := "192.168." + m + "." + tt

	value := startIP + " " + stopIP + ";"

	modifyfile.ChangeFileParameter(DHCP_CONF, "range", value)
}

func SetRentTime(time string) {
	modifyfile.ChangeFileParameter(DHCP_CONF, "default-lease-time", time)
}

func SetDHCPDNS(twoDNS string) {
	modifyfile.ChangeFileParameter(DHCP_CONF, "option domain-name-servers", twoDNS)
}
