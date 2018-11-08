package netconfig

import (
	"fmt"
	"io/ioutil"
	"modifyfile"
	"net"
	"os"
	"os/exec"
	"strings"
)

const RESOLV_CONF = "/etc/resolv.conf"

type IPConfiguration struct {

	//is port name default eth0 ?
	PortName string
	IP []string
	Mask []string
	GW string
	MTU string
	PrimaryDNS string
	OptionalDNS string
}

//change MTU by using command line "ifconfig"
func (ipConf *IPConfiguration) SetMTU() {
	cmd := exec.Command("ifconfig", ipConf.PortName, "mtu", ipConf.MTU)
	cmd.Run()
}

//change the IP directly by using command line "ifconfig"
func (ipConf *IPConfiguration) SetIP() {
	var newIP = ipConf.IP[0] + "." + ipConf.IP[1] + "." + ipConf.IP[2] + "." + ipConf.IP[3]
	var newMask = ipConf.Mask[0] + "." + ipConf.Mask[1] + "." + ipConf.Mask[2] + "." + ipConf.Mask[3]
	cmd := exec.Command("ifconfig", ipConf.PortName, newIP, "netmask", newMask)
	cmd.Run()
}

//change the gateway directly by using command line "route"
func (ipConf *IPConfiguration) SetGW() {
	cmd := exec.Command("route",  "add", "default", "gateway", ipConf.GW)
	cmd.Run()
}


func (ipConf *IPConfiguration) SetPrimaryDNS() {
	modifyfile.ChangeFileParameter(RESOLV_CONF, "nameserver", ipConf.PrimaryDNS)
}

func (ipConf *IPConfiguration) SetOptionalDNS() {
	b, err := ioutil.ReadFile(RESOLV_CONF)
	if err != nil {
		fmt.Println(err)
	}

	DNS1 := modifyfile.FindParameterLocation(b, "nameserver")
	for ; b[DNS1] != '\n'; {
		DNS1++
	}
	DNS2 := DNS1 + 12

	b1 := append([]byte{}, b[:DNS2]...)
	b2 := []byte(ipConf.OptionalDNS)

	d := append(b1, b2...)
	ioutil.WriteFile(RESOLV_CONF, d, 0777)
}

//get IP and Mask information, it returns two string slices
//for example: {"192", "168", "10", "1"} and {"255", "255", "255", "0"}
func (ipConf *IPConfiguration) GetIPandMask() {
	var b []byte
	var m []byte

	cmd := exec.Command("ifconfig", ipConf.PortName)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	l1 := strings.Index(string(stdout), "inet addr") + 10
	for i := l1; stdout[i] != ' '; i++ {
		b = append(b, stdout[i])
	}
	bb := strings.Split(string(b), ".")

	l2 := strings.Index(string(stdout), "Mask") + 5
	for i := l2; stdout[i] != '\n'; i++ {
		m = append(m, stdout[i])
	}
	mm := strings.Split(string(m), ".")

	ipConf.IP = bb
	ipConf.Mask = mm
}

//get default gateway information by command "route -n"
//it returns a string slice, for example: {"192", "168", "1", "1"}
func GetGateway() []string {
	var g []byte

	cmd := exec.Command("route", "-n")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	l := strings.Index(string(stdout), "0.0.0.0") + 16
	for i := l; stdout[i] != ' '; i++ {
		g = append(g, stdout[i])
	}
	gg := strings.Split(string(g), ".")

	return gg
}

//get DNS server address by read file "/etc/resolv.conf"
//it returns a string slice, for example: {"8", "8", "8", "8"}
//if there are two DNS server, merging them into one slice string, for example: {"8", "8", "8", "8", "8", "8", "4", "4"}
func GetDNS() []string {
	var n []byte

	b, err := ioutil.ReadFile(RESOLV_CONF)
	if err != nil {
		fmt.Println(err)
	}

	str := string(b)
	l1 := strings.Index(str, "nameserver") + len("nameserver") + 1
	for i := l1; b[i] != '\n'; i++ {
		n = append(n, b[i])
	}

	b2 := append([]byte{}, b[l1+len(n):]...)
	l2 := strings.Index(string(b2), "nameserver")
	if l2 != -1 {
		n = append(n, '.')
		for i := l2 + 11; b2[i] != '\n'; i++ {
			n = append(n, b2[i])
		}
	}

	nn := strings.Split(string(n), ".")
	return nn
}

func GetDomainName() []byte {
	return modifyfile.ReadFileParameter(RESOLV_CONF, "domain")
}

//write ip configuration to profile(path: /etc/network/interfaces), only reboot will work
func (ipConf *IPConfiguration) WriteIP() {
	file, err := os.OpenFile("/home/hyh/Templates/testfile", os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var newIP = ipConf.IP[0] + "." + ipConf.IP[1] + "." + ipConf.IP[2] + "." + ipConf.IP[3]
	var newMask = ipConf.Mask[0] + "." + ipConf.Mask[1] + "." + ipConf.Mask[2] + "." + ipConf.Mask[3]
	n, err := file.WriteString("\n" + "auto " + ipConf.PortName + "\n" + "iface " + ipConf.PortName + " inet static\n" + "address " + newIP + "\n" + "netmask " + newMask + "\n" + "gateway " + ipConf.GW)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func (ipConf *IPConfiguration) show() {

	//cannot get [ip and mask and gateway] by this way
	eth, err := net.Interfaces()
	if err != nil {
		fmt.Println(eth)
	}
	fmt.Println(eth)
	for _, i := range eth {
		fmt.Println("********************")
		fmt.Println("Index:", i.Index)
		fmt.Println("MTU:", i.MTU)
		fmt.Println("Name:", i.Name)
		fmt.Println("HardwareAddr:", i.HardwareAddr)
		fmt.Println("Flags:", i.Flags)
	}
	fmt.Println()


	//show interface information after execute "ifconfig"
	cmd := exec.Command("ifconfig")
	stdout, err := cmd.Output()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%s", stdout)
	fmt.Println()

	//get [ip mask] information
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, address := range addrs {
		fmt.Println("************")
		fmt.Println("netname:", address.Network())
		fmt.Println("netaddress:", address.String())
	}
}
