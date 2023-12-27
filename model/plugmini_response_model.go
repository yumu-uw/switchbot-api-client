package model

type PlugMiniResponseModel struct {
	BaseResponseModel
	Device PlugMini `json:"body"`
}

type PlugMini struct {
	DeviceID         string  `json:"deviceId"`
	DeviceType       string  `json:"deviceType"`
	HubDeviceID      string  `json:"hubDeviceId"`
	Power            string  `json:"power"`
	Voltage          float64 `json:"voltage"`
	Weight           float64 `json:"weight"`
	ElectricityOfDay int     `json:"electricityOfDay"`
	ElectricCurrent  float64 `json:"electricCurrent"`
	Version          string  `json:"version"`
}
