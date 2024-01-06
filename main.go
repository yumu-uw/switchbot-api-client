package main

import (
	"github.com/joho/godotenv"
	"github.com/yumu-uw/switchbot-api-client/exporter"
	"github.com/yumu-uw/switchbot-api-client/util"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("can't read .env file")
	}
}

func main() {
	loadEnv()
	util.InitApiUtil()

	var status_list []map[string]interface{}
	device_list := util.GetDeviceList()
	for _, d := range device_list {
		r := util.GetDeviceStatus(d.DeviceID)
		r["DeviceName"] = d.DeviceName
		r["DeviceType"] = d.DeviceType
		status_list = append(status_list, r)
	}

	for _, v := range status_list {
		device_type, _ := v["DeviceType"].(string)
		device_name, _ := v["DeviceName"].(string)
		body := v["body"].(map[string]interface{})

		tags := map[string]string{
			"device_type": device_type,
			"device_name": device_name,
		}
		exporter.InitInfluxDBExporter()
		exporter.ExportToInfluxDB(tags, body)
	}
	// bytes, err := json.Marshal(o)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(bytes))
	// util.DBTest()
}
