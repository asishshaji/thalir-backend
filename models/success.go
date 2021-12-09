package models

type ResponseSuccess struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
}
