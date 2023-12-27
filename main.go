package main

import (
	"github.com/yumu-uw/switchbot-api-client/util"
)

func main() {
	device_list := util.GetDeviceList()
	for _, d := range device_list {
		switch d.DeviceType {
		case "Plug Mini (JP)", "Plug Mini (US)":
			util.GetPlugMiniStatus(d.DeviceID)
		}
	}
}
