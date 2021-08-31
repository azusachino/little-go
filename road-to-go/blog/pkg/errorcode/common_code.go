package errorcode

var (
	Success              = NewError(0, "success")
	ServerError          = NewError(10000, "server internal error")
	InvalidParams        = NewError(10001, "invalid param")
	NotFound             = NewError(10002, "not found")
	AuthNotExist         = NewError(10003, "auth not exist")
	TokenError           = NewError(10004, "token error")
	TokenTimeout         = NewError(10005, "token timeout")
	TokenGenerateFailure = NewError(10006, "token generate failed")
	TooManyRequests      = NewError(10007, "too many requests")
)
