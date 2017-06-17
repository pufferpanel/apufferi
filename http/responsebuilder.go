package http

type responseBuilder struct {
	response response
	context  *Context
}

type Builder interface {
	Send()

	Code(code int) Builder
	WithCode(code int) Builder

	Message(message string) Builder
	WithMessage(message string) Builder

	Data(data interface{}) Builder
	WithData(data interface{}) Builder

	MessageCode(code Code) Builder
	WithMessageCode(code Code) Builder

	Fail() Builder
	Success() Builder
	WithSuccess(success bool) Builder
}


func Respond(c *Context) Builder {
	return responseBuilder{
		response: response{
			Success:     true,
			Status:        200,
			Code: SUCCESS,
		},
		context: c,
	}
}

func (rb responseBuilder) Status(code int) Builder {
	return rb.WithStatus(code)
}

func (rb responseBuilder) WithStatus(code int) Builder {
	rb.response.Code = code
	if code > 299 || code < 200 {
		return rb.Fail()
	}
	return rb
}

func (rb responseBuilder) Message(message string) Builder {
	return rb.WithMessage(message)
}

func (rb responseBuilder) WithMessage(message string) Builder {
	rb.response.Message = message
	return rb
}

func (rb responseBuilder) Data(data interface{}) Builder {
	return rb.WithData(data)
}

func (rb responseBuilder) WithData(data interface{}) Builder {
	rb.response.Data = data
	return rb
}

func (rb responseBuilder) Code(code Code) Builder {
	return rb.WithCode(code)
}

func (rb responseBuilder) WithCode(code Code) Builder {
	rb.response.Code = code
	return rb
}

func (rb responseBuilder) Fail() Builder {
	return rb.WithSuccess(false)
}

func (rb responseBuilder) Success() Builder {
	return rb.WithSuccess(true)
}

func (rb responseBuilder) WithSuccess(success bool) Builder {
	rb.response.Success = success
	return rb
}

func (rb responseBuilder) Send() {
	rb.context.JSON(rb.response.Code, rb.response)
}
