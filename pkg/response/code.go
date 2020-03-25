package response

const (
	SUCCESS 						= 0
	SERVER_ERROR 					= 1000 // 系统错误
	NOT_FOUND       				= 1001 // 401错误
	ERROR_PARAMS 					= 1002  // 参数错误
	ERROR_AUTHORIZATION             = 1003 // 未认证
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT 	= 20002 // token超时
	ERROR_AUTH_TOKEN 				= 20003 // token生成错误
	ERROR_AUTH 						= 20004 // token认证错误
)