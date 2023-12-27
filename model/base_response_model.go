package model

type BaseResponseModel struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
