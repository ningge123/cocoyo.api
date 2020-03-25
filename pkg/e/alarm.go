package e

import (
	"cocoyo/pkg/util"
	"encoding/json"
	"log"
	"runtime/debug"
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

type errorInfo struct {
	Time 		string 	`json:"time"`
	Alarm 		string 	`json:"alarm"`
	Message 	string 	`json:"message"`
	Stack       string  `json:"stack"`
}

func New(level, text string) {
	alarm(level, text)
}

func alarm(level string, str string) {

	debugStack := ""
	for _, v := range strings.Split(string(debug.Stack()), "\n") {
		debugStack += v + " "
	}

	var msg = errorInfo {
		Time: util.GetTimeStr(),
		Alarm: level,
		Message: str,
		Stack: debugStack,
	}

	jsons, _ := json.Marshal(msg)

	errorJsonInfo := string(jsons)

	if level == ERROR {
		log.Printf("[%s] %s %s", level, str, errorJsonInfo)

		return
	}

	log.Printf("[%s] %s %s", level, str, errorJsonInfo)
}