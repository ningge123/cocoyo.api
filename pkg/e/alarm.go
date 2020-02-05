package e

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

// 日志等级
const (
	DEBUG     = "DEBUG"
	INFO      = "INFO"
	NOTICE    = "NOTICE"
	WARNING   = "WARNING"
	ERROR     = "ERROR"
	CRITICAL  = "CRITICAL"
	ALERT     = "ALERT"
	EMERGENCY = "EMERGENCY"
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

func New(level, text string) error {
	alarm(level, text)

	return &errorString{text}
}

func alarm(level string, str string) {
	// 定义 文件名、行号、方法名
	fileName, line, functionName := "?", 0 , "?"

	pc, fileName, line, ok := runtime.Caller(2)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}

	var msg = errorInfo {
		Filename : fileName,
		Line     : line,
		Funcname : functionName,
	}

	jsons, _ := json.Marshal(msg)

	errorJsonInfo := string(jsons)

	if level == ERROR {
		log.Fatal(fmt.Sprintf("[%s] %s %s", level, str, errorJsonInfo))
	}

	log.Println(fmt.Sprintf("[%s] %s %s", level, str, errorJsonInfo))
}