package model

type BaseStatusModel struct {
	DeviceID   string `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	DeviceType string `json:"deviceType"`
	Version    string `json:"version"`
}
