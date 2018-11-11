package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"restapi"
	"strconv"
)

type TopoController struct {
	beego.Controller
}

type Topology struct {
	Result []Contex `json:"data"`
}
type Contex struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Sj string `json:"sj"`
	ID string `json:"-"`
}

func (c *TopoController) Get() {
	topo := topoAssemble()
	fmt.Println(topo)
	b, err := json.Marshal(topo)
	for _, v := range b {
		fmt.Printf("%c", v)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, string(b))
}


func topoAssemble() Topology {
	layerOne := theFirstLayer()
	layerTwo := theSecondLayer()
	layerThree := theThirdLayer()

/***************************set value for data in layerThree according layerTwo**********************************/
	// use a map(key: organizationID, value: orgData) to assignment devData.Sj
	var m1 = make(map[string]Contex)
	for _, v := range layerTwo {
		m1[v.ID] = v
	}
	for i := 0; i < len(layerThree); i++ {
		if v, ok := m1[layerThree[i].ID]; ok {
			layerThree[i].Sj = v.Value
		}
	}

	// use a map(key: Sj, value: Value) to assignment devData.Value
	var m2 = make(map[string]string)
	for i := 0; i < len(layerThree); i++ {
		if v, ok := m2[layerThree[i].Sj]; ok {
			newValue := increaseValue(v)
			layerThree[i].Value = newValue
			m2[layerThree[i].Sj] = newValue
		}else {
			layerThree[i].Value = layerThree[i].Sj + "01"
			m2[layerThree[i].Sj] = layerThree[i].Sj + "01"
		}
	}

/************************************assemble all the data**************************************************/
	var t Topology
	t.Result = append(t.Result, layerOne)
	t.Result = append(t.Result, layerTwo...)
	t.Result = append(t.Result, layerThree...)
	return t
}

// increase the value of devData.Value by 1
func increaseValue(s string) string {
	var b []byte
	for i := 1; i < len(s); i++ {
		b = append(b, s[i])
	}
	number, _ := strconv.Atoi(string(b))
	number++
	return "0" + strconv.Itoa(number)
}

// the first layer of topology is server
func theFirstLayer() Contex {
	var serverData Contex
	servers := restapi.GetServers("100")
	serverData.Name = servers.Result[0].Name
	serverData.Value = "01"
	serverData.Sj = "-"

	return serverData
}

// the second layer of topology is organization
func theSecondLayer() []Contex {
	var orgData []Contex
	orgs := restapi.GetOrganizations("100")
	for _, v := range orgs.Result {
		orgData = append(orgData, Contex{Name: v.Name, ID: v.ID})
	}
	for i := 0; i < len(orgData); i++ {
		orgData[i].Sj = "01"
		val := "01" + "0" + strconv.Itoa(i+1)
		orgData[i].Value = val
	}

	return orgData
}

// the third layer of topology is device
func theThirdLayer() []Contex {
	var devData []Contex

	// get a Result list of device which is tagged with organization ID
	// devs is like this: { [Result1:{[deviceName, orgID] [deviceName, orgID]}]     [Result2: {}]     [Result3: {}] }
	var devs []restapi.DeviceWithOrgIDList
	apps := restapi.GetAllApplications("100")
	for _, v := range apps.Result {
		d := restapi.GetDevicesWithOrgID(v.ID, v.OrganizationID, "100")
		devs = append(devs, *d)
	}

	// use double cycle to parse devs to one by one Contex structure
	for _, result := range devs {
		for _, v := range result.Result {
			devData = append(devData, Contex{Name: v.Name, ID: v.OrgID})
		}
	}

	return devData
}