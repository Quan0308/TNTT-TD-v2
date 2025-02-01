package dtos

type Response struct {
	StatusCode int         `json:"-"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
