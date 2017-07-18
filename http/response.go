package http

type response struct {
	Success     bool        `json:"success"`
	Message     string      `json:"msg,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Status      int         `json:"-"`
	Code        Code 	`json:"code,omitempty"`
}
