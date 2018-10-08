package netconfig

import (
	"os"
	"fmt"
	"io/ioutil"
)

type NetTransStatistics struct {
	RxPackets string
	RxErrors string
	RxDropped string
	RxBytes string

	TxPackets string
	TxErrors string
	TxDropped string
	TxCollisions string
	TxBytes string
}

func (eth *NetTransStatistics) GetEthSta() {

	//RX INFORMATION
	f1, err := os.Open("/sys/class/net/eth0/statistics/rx_packets")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()
	cont1, _ := ioutil.ReadAll(f1)
	eth.RxPackets = string(cont1)

	f2, err := os.Open("/sys/class/net/eth0/statistics/rx_errors")
	if err != nil {
		fmt.Println(err)
	}
	defer f2.Close()
	cont2, _ := ioutil.ReadAll(f2)
	eth.RxErrors = string(cont2)

	f3, err := os.Open("/sys/class/net/eth0/statistics/rx_dropped")
	if err != nil {
		fmt.Println(err)
	}
	defer f3.Close()
	cont3, _ := ioutil.ReadAll(f3)
	eth.RxDropped = string(cont3)

	f5, err := os.Open("/sys/class/net/eth0/statistics/rx_bytes")
	if err != nil {
		fmt.Println(err)
	}
	defer f5.Close()
	cont5, _ := ioutil.ReadAll(f5)
	eth.RxBytes = string(cont5)


	//TX INFORMATION
	f6, err := os.Open("/sys/class/net/eth0/statistics/tx_packets")
	if err != nil {
		fmt.Println(err)
	}
	defer f6.Close()
	cont6, _ := ioutil.ReadAll(f6)
	eth.TxPackets = string(cont6)

	f7, err := os.Open("/sys/class/net/eth0/statistics/tx_errors")
	if err != nil {
		fmt.Println(err)
	}
	defer f7.Close()
	cont7, _ := ioutil.ReadAll(f7)
	eth.TxErrors = string(cont7)

	f8, err := os.Open("/sys/class/net/eth0/statistics/tx_dropped")
	if err != nil {
		fmt.Println(err)
	}
	defer f8.Close()
	cont8, _ := ioutil.ReadAll(f8)
	eth.TxDropped = string(cont8)

	f10, err := os.Open("/sys/class/net/eth0/statistics/tx_bytes")
	if err != nil {
		fmt.Println(err)
	}
	defer f10.Close()
	cont10, _ := ioutil.ReadAll(f10)
	eth.TxBytes = string(cont10)

	f11, err := os.Open("/sys/class/net/eth0/statistics/collisions")
	if err != nil {
		fmt.Println(err)
	}
	defer f11.Close()
	cont11, _ := ioutil.ReadAll(f11)
	eth.TxCollisions = string(cont11)
}


func (wlan *NetTransStatistics) GetWlanSta() {

	//RX INFORMATION
	f1, err := os.Open("/sys/class/net/wlan0/statistics/rx_packets")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()
	cont1, _ := ioutil.ReadAll(f1)
	wlan.RxPackets = string(cont1)

	f2, err := os.Open("/sys/class/net/wlan0/statistics/rx_errors")
	if err != nil {
		fmt.Println(err)
	}
	defer f2.Close()
	cont2, _ := ioutil.ReadAll(f2)
	wlan.RxErrors = string(cont2)

	f3, err := os.Open("/sys/class/net/wlan0/statistics/rx_dropped")
	if err != nil {
		fmt.Println(err)
	}
	defer f3.Close()
	cont3, _ := ioutil.ReadAll(f3)
	wlan.RxDropped = string(cont3)

	f5, err := os.Open("/sys/class/net/wlan0/statistics/rx_bytes")
	if err != nil {
		fmt.Println(err)
	}
	defer f5.Close()
	cont5, _ := ioutil.ReadAll(f5)
	wlan.RxBytes = string(cont5)


	//TX INFORMATION
	f6, err := os.Open("/sys/class/net/wlan0/statistics/tx_packets")
	if err != nil {
		fmt.Println(err)
	}
	defer f6.Close()
	cont6, _ := ioutil.ReadAll(f6)
	wlan.TxPackets = string(cont6)

	f7, err := os.Open("/sys/class/net/wlan0/statistics/tx_errors")
	if err != nil {
		fmt.Println(err)
	}
	defer f7.Close()
	cont7, _ := ioutil.ReadAll(f7)
	wlan.TxErrors = string(cont7)

	f8, err := os.Open("/sys/class/net/wlan0/statistics/tx_dropped")
	if err != nil {
		fmt.Println(err)
	}
	defer f8.Close()
	cont8, _ := ioutil.ReadAll(f8)
	wlan.TxDropped = string(cont8)

	f10, err := os.Open("/sys/class/net/wlan0/statistics/tx_bytes")
	if err != nil {
		fmt.Println(err)
	}
	defer f10.Close()
	cont10, _ := ioutil.ReadAll(f10)
	wlan.TxBytes = string(cont10)

	f11, err := os.Open("/sys/class/net/wlan0/statistics/collisions")
	if err != nil {
		fmt.Println(err)
	}
	defer f11.Close()
	cont11, _ := ioutil.ReadAll(f11)
	wlan.TxCollisions = string(cont11)
}
