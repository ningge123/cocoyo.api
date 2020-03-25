package response

type Api struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

var data interface{}

func ServerError() *Api {
	return newApiResponse(SERVER_ERROR, GetMsg(SERVER_ERROR), data)
}

func NotFound() *Api {
	return newApiResponse(NOT_FOUND, GetMsg(NOT_FOUND), data)
}

func Authorization() *Api {
	return newApiResponse(ERROR_AUTHORIZATION, GetMsg(ERROR_AUTHORIZATION), data)
}

func ParameterError(message string) *Api {
	return newApiResponse(ERROR_PARAMS, message, data)
}

func Response(data interface{}) *Api {
	return newApiResponse(SUCCESS, GetMsg(SUCCESS), data)
}

func SuccessNotContent() *Api {
	return newApiResponse(SUCCESS, GetMsg(SUCCESS), data)
}

func newApiResponse(code int, message string, data interface{}) *Api {
	return &Api{
		Code: code,
		Message: message,
		Data: data,
	}
}