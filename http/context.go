package http

type Context interface {
	JSON(code int, body interface{})
}
