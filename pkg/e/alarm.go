package e

import (
	"cocoyo/pkg/function"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type errorString struct {
	s string
}

type errorInfo struct {
	Time 		string 	`json:"time"`
	Alarm 		string 	`json:"alarm"`
	Message 	string 	`json:"message"`
	Filename 	string 	`json:"filename"`
	Line 		int 	`json:"line"`
	Funcname 	string 	`json:"funcname"`
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	alarm("INFO", text)

	return &errorString{text}
}

func alarm(level string, str string) {
	// 当前时间
	currentTime := function.GetTimeStr()

	// 定义 文件名、行号、方法名
	fileName, line, functionName := "?", 0 , "?"

	pc, fileName, line, ok := runtime.Caller(2)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}

	var msg = errorInfo {
		Time     : currentTime,
		Alarm    : level,
		Message  : str,
		Filename : fileName,
		Line     : line,
		Funcname : functionName,
	}

	jsons, errs := json.Marshal(msg)

	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}

	errorJsonInfo := string(jsons)

	fmt.Println(errorJsonInfo)
}