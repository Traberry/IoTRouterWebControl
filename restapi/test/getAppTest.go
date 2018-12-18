package main

import (
	"fmt"
	"restapi"
)

func main() {
	var s *restapi.DeviceActivation

	s = restapi.GetDeviceActivation("bbbbbbbbbbbbbbbb")

	fmt.Println(s)
}
