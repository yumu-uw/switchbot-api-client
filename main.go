package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yumu-uw/switchbot-api-client/util"
)

func main() {
	o := make(map[string][]interface{})
	device_list := util.GetDeviceList()
	for _, d := range device_list {
		switch d.DeviceType {
		case "Plug Mini (JP)", "Plug Mini (US)":
			r := util.GetPlugMiniStatus(d.DeviceID)
			r.Device.DeviceName = d.DeviceName
			o[d.DeviceType] = append(o[d.DeviceType], r)
		case "Hub 2":
			r := util.GetHub2Status(d.DeviceID)
			r.Device.DeviceName = d.DeviceName
			o[d.DeviceType] = append(o[d.DeviceType], r)
		}
	}
	bytes, err := json.Marshal(o)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytes))
}
