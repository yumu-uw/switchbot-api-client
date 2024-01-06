package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/yumu-uw/switchbot-api-client/exporter"
	"github.com/yumu-uw/switchbot-api-client/util"
)

var output string
var outputPath string

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("can't read .env file")
	}
}

func parseArgs() {
	flag.StringVar(&output, "output", "json", "取得したデータの出力方法")
	flag.StringVar(&outputPath, "path", "./output.json", "jsonの出力先（outputがjsonの場合のみ有効）")
	flag.Parse()
}

func main() {
	loadEnv()
	parseArgs()
	util.InitApiUtil()

	var status_list []map[string]interface{}
	device_list := util.GetDeviceList()
	for _, d := range device_list {
		r := util.GetDeviceStatus(d.DeviceID)
		r["DeviceName"] = d.DeviceName
		r["DeviceType"] = d.DeviceType
		status_list = append(status_list, r)
	}

	switch output {
	case "json":
		exporter.ExportToJson(status_list)
	case "influxdb":
		exporter.InitInfluxDBExporter()
		for _, v := range status_list {
			device_type, _ := v["DeviceType"].(string)
			device_name, _ := v["DeviceName"].(string)
			body := v["body"].(map[string]interface{})

			tags := map[string]string{
				"device_type": device_type,
				"device_name": device_name,
			}
			exporter.ExportToInfluxDB(tags, body)
		}
	}
}
