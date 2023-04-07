package domain

type ErrorFormat struct {
	Code    int
	Message string
}

type ResponseFormat struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
}

var (
	ErrorForbidden           = ErrorFormat{Code: 403, Message: "Forbidden"}
	ErrorAuthTokenExpired    = ErrorFormat{Code: 4011, Message: "auth token expired"}
	ErrorInvalidToken        = ErrorFormat{Code: 4012, Message: "invalid token"}
	ErrorPermission          = ErrorFormat{Code: 4013, Message: "Permission denied"}
	ErrorBadRequest          = ErrorFormat{Code: 400, Message: "bad request"}
	ErrorServer              = ErrorFormat{Code: 500, Message: "Server Error"}
	ErrorUnknowInternalError = ErrorFormat{Code: 404, Message: "Unknow Internal Error"}
)
