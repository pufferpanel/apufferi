package http

type Code uint

const (
	SUCCESS          Code = 0
	NOTAUTHORIZED         = 400
	NOAUTHENTICATION      = 401
	NOSERVER              = 402
	MALFORMEDJSON         = 403
	NOFILE                = 404
	NOSERVERID            = 405
	INVALIDTIME           = 406
	INVALIDREQUEST        = 407
	UNKNOWN               = 999
)
