package response

var MsgFlags = map[int]string{
	SUCCESS:                        "success",
	SERVER_ERROR:                   "系统错误",
	ERROR_PARAMS:                   "请求参数错误",
	NOT_FOUND: 						"404",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_AUTHORIZATION:            "Unauthorized.",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]

	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}