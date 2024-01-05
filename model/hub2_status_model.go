package model

type Hub2ResponseModel struct {
	BaseResponseModel
	Device Hub2 `json:"body"`
}

type Hub2 struct {
	BaseStatusModel
	Temperature float32 `json:"temperature"`
	LightLevel  int     `json:"lightLevel"`
	Humidity    int     `json:"humidity"`
}
