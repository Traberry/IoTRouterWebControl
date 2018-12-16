package restapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJpc3MiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJuYmYiOjE1MzU0NDE5OTUsInN1YiI6InVzZXIiLCJ1c2VybmFtZSI6ImFkbWluIn0.-agKC7UzJL6fOordtJb7qVBwbGeUa8TW__C0LHlniXw"
	//anotherToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJpc3MiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJuYmYiOjE1MzY3MTgzMDAsInN1YiI6InVzZXIiLCJ1c2VybmFtZSI6ImFkbWluIn0.wYaXH-XmpzoIwsSxCW1shPwCZ7F4x0PM0n0xmOzhiU8"
	//anotherIPAddress = "121.49.107.48:8080"
	IPAddressOfAPIServer = "192.168.3.105:8080"
)

type Application struct {
	ID string
	Name string
	Description string
	OrganizationID string
	//ServiceProfileID string
	ServiceProfileName string
}

type Gateway struct {
	MAC string
	Name string
	//Description string
	//CreatedAt string
	//UpdatedAt string
	//OrganizationID string
	//NetworkServerID string
	Activity string
}

type GatewayDtails struct {
	MAC string
	Name string
	//Description string
	Latitude int
	Longtitude int
	Altitude int
	//CreatedAt string
	//UpdatedAt string
	//FirstSeenAt string
	LastSeenAt string
	//OrganizationID string
	//Ping bool
	//NetworkServerID string
	//GatewayProfileID string
}

type Device struct {
	DevEUI string
	Name string
	//ApplicationID string
	Description string
	//DeviceProfileID string
	DeviceProfileName string
	DeviceStatusBattery int
	DeviceStatusMargin int
	LastSeenAt string
}

type SimpleDevice struct {
	Name string
	DeviceStatusBattery int
	DeviceStatusMargin int
}

type DeviceWithOrgID struct {
	Name string
	OrgID string
}

type Organization struct {
	ID string
	Name string
	DisplayName string
	CanHaveGateways bool
	//CreatedAt string
	//UpdatedAt string
}

type User struct {
	ID string
	UserName string
	//SessionTTL int
	IsAdmin bool
	IsActive bool
	//CreatedAt string
	//UpdatedAt string
	//Email string
	//Note string
}

type NetworkServer struct {
	//ID string
	//CreatedAt string
	//UpdatedAt string
	Name string
	Server string
	//CaCert string
	//TLSCert string
	//RoutingProfileCACert string
	//RoutingProfileTLSCert string
	//GatewayDiscoveryEnabled bool
	//GatewayDiscoveryInterval int
	//GatewayDiscoveryTxFrequency int
	//GatewayDiscoveryDR int
	//Version string
	//Region string
}

type ApplicationList struct {
	TotalCount string
	Result []Application
}

type GatewayList struct {
	TotalCount int
	Result []Gateway
}

type DeviceList struct {
	//TotalCount string
	Result []Device
}

type SimpleDeviceList struct {
	Result []SimpleDevice
}

type DeviceWithOrgIDList struct {
	Result []DeviceWithOrgID
}

type OrganizationList struct {
	//TotalCount int
	Result []Organization
}

type UserList struct {
	//TotalCount int
	Result []User
}

type NetworkServerList struct {
	//TotalCount string
	Result []NetworkServer
}

func MakeAPIRequest(url, method string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify:true},
	}
	client := &http.Client{Transport:tr}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logs.Warn("make a bad API request: %v", err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	//req.Header.Add("Grpc-Metadata-Authorization", token)
	req.Header.Add("Grpc-Metadata-Authorization", token)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//Would it be better if return a pointer instead of a value ???
//I changed the return from a value to a pointer, but I have not tried it yet (2018-09-04)
func GetAllApplications(limit string) *ApplicationList {
	var s ApplicationList
	url := "https://" + IPAddressOfAPIServer + "/api/applications?limit=" + limit
	//url := "https://" + anotherIPAddress + "/api/applications?limit=" + limit

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetAllApplications: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetAllApplications response is NOT OK: %v", response.StatusCode)
		return nil
	}

	return &s
}

func GetAllGateways(limit string) *GatewayList {
	var s GatewayList
	url := "https://" + IPAddressOfAPIServer + "/api/gateways?limit=" + limit

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetAllGateways: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetAllGateways response is NOT OK: %v", response.StatusCode)
		return nil
	}

	return &s
}

func GetGatewayDetails(macAddress string) *GatewayDtails {
	var s GatewayDtails
	url := "https://" + IPAddressOfAPIServer + "/api/gateways/" + macAddress

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetGatewayDetails: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetGatewayDetails response is NOT OK: %v", response.StatusCode)
		return nil
	}

	return &s
}

func GetDevices(appID, limit string) *DeviceList {
	var s DeviceList
	url := "https://" + IPAddressOfAPIServer + "/api/applications/" + appID + "/devices?limit=" + limit

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetDevices: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetDevices response is NOT OK: %v", response.StatusCode)
		return nil
	}

	return &s
}

func GetSimpleDevices(appID, limit string) *SimpleDeviceList {
	var s SimpleDeviceList
	url := "https://" + IPAddressOfAPIServer + "/api/applications/" + appID + "/devices?limit=" + limit

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetSimpleDevices: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetSimpleDevices response is NOT OK: %v", response.StatusCode)
		return nil
	}

	return &s
}

func GetDevicesWithOrgID(appID, organizationID, limit string) *DeviceWithOrgIDList {
	var s DeviceWithOrgIDList
	url := "https://" + IPAddressOfAPIServer + "/api/applications/" + appID + "/devices?limit=" + limit

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetSimpleDevices: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetSimpleDevices response is NOT OK: %v", response.StatusCode)
		return nil
	}

	for i := 0; i < len(s.Result); i++ {
		s.Result[i].OrgID = organizationID
	}

	return &s
}

func GetOrganizations(limit string) *OrganizationList {
	var s OrganizationList
	url := "https://" + IPAddressOfAPIServer + "/api/organizations?limit=" + limit

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetOrganizations: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetOrganizations response is NOT OK: %v", response.StatusCode)
		return nil
	}

	return &s
}

func GetUsers(limit string) *UserList {
	var s UserList
	url := "https://" + IPAddressOfAPIServer + "/api/users?limit=" + limit

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetUsers: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetUsers response is NOT OK: %v", response.StatusCode)
		return nil
	}

	return &s
}

func GetServers(limit string) *NetworkServerList {
	var s NetworkServerList
	url := "https://" + IPAddressOfAPIServer + "/api/network-servers?limit=" + limit

	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetServers: %v", err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		logs.Warn("GetServers response is NOT OK: %v", response.StatusCode)
		return nil
	}

	return &s
}

type GatewaySta struct {
	//Timestamp string
	RxPacketsReceived int
	//RxPacketsReceivedOK int
	//TxPacketsReceived int
	//TxPacketsEmitted int
}

type GatewayStates struct {
	Result []GatewaySta
}

func GetGatewayActivity(mac string) string {
	var activity string
	var s GatewayStates
	url := "https://" + IPAddressOfAPIServer + "/" + mac + "/stats?interval=day&startTimestamp=" + "2018-11-18T15:04:05.999999999Z" + "&endTimestamp=" + "2018-12-15T15:04:05.999999999Z"
	fmt.Println(url)
	response, err := MakeAPIRequest(url, "GET")
	if err != nil {
		logs.Error("Error while making API request to GetServers: %v", err.Error())
		return ""
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
		for _, v := range s.Result {
			activity += strconv.Itoa(v.RxPacketsReceived)
			activity += ","
		}
	}else {
		logs.Warn("GetGatewayActivity response is NOT OK: %v", response.StatusCode)
		return ""
	}
	return activity
}


/*
func GetGatewayState(mac string) *GatewayStates {
	var s GatewayStates
	url := "https://192.168.3.109:8080/api/gateways/" + mac + "/stats?interval=day&startTimestamp=" + "2018-08-29T15:04:05.999999999Z" + "&endTimestamp=" + "2018-09-27T15:04:05.999999999Z"
	response := MakeAPIRequest(url, "GET")
	defer response.Body.Close()
	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &s)
	}else {
		fmt.Println(response.Status)
	}
	return &s
}
*/

/*
func DeviceListByAppID() {
	var appID []string
	appList := GetAllApplications("10")
	count, _ := strconv.Atoi(appList.TotalCount)
	for i := 0; i < count; i++ {
		appID = append(appID, appList.Result[i].ID)
	}

	for _, v := range appID {
		fmt.Printf("*****Application %s*****\n", v)
		deviceList := GetDevices( v, "10")
		deviceCount, _ := strconv.Atoi(deviceList.TotalCount)
		for i := 0; i < deviceCount; i++ {
			fmt.Println(deviceList.Result[i])
		}
	}
}*/
