package main

import (
	"encoding/json"
	"fmt"
	"restapi"
)

type GWs struct {
	Result []restapi.Gateway
}

func main() {
	gw := restapi.GetAllGateways("10")
	//fmt.Println(gw)

	for i := 0; i < gw.TotalCount; i++ {
		gw.Result[i].Activity = restapi.GetGatewayActivity(gw.Result[i].MAC)
	}
	//fmt.Println(gw)

	var gateways GWs
	gateways.Result = gw.Result
	b, err := json.Marshal(gateways)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range b {
		fmt.Printf("%c", v)
	}
}
