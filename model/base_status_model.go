package model

import (
	"bytes"
	"encoding/json"
)

type BaseStatusModel struct {
	DeviceID    string `json:"deviceId"`
	DeviceName  string `json:"deviceName"`
	DeviceType  string `json:"deviceType"`
	Version     string `json:"version"`
	HubDeviceID string `json:"hubDeviceId"`
}

func (b BaseStatusModel) ConvertToMap(data interface{}) (map[string]interface{}, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	out := new(bytes.Buffer)
	// JSONの整形
	err = json.Indent(out, jsonStr, "", "    ")
	if err != nil {
		return nil, err
	}
	var mapData map[string]interface{}
	if err := json.Unmarshal(out.Bytes(), &mapData); err != nil {
		return nil, err
	}
	return mapData, err
}
