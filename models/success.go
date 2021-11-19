package models

type ResponseSuccess struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
