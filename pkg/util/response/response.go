package response


type Code struct {
	Code 	int
	Message string
}

var Code200 = Code{Code: SUCCESS, Message: "success"}

type Response struct {
	Code 	int `json:"code"`
	Message string `json:"message"`
	Data 	interface{} `json:"data"`
}


func ReturnMsgFunc(code Code,data interface{}) *Response {
	rm := &Response{Code: code.Code, Message: code.Message, Data: data}

	return rm
}