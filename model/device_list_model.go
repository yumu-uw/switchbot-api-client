package model

type DeviceListResponseModel struct {
	BaseResponseModel
	Body Body `json:"body"`
}

type Body struct {
	DeviceList         []Device `json:"deviceList"`
	InfraredRemoteList []any    `json:"infraredRemoteList"`
}

type Device struct {
	DeviceID           string `json:"deviceId"`
	DeviceName         string `json:"deviceName"`
	DeviceType         string `json:"deviceType"`
	EnableCloudService bool   `json:"enableCloudService"`
	HubDeviceID        string `json:"hubDeviceId"`
}
