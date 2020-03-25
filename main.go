package main

import (
	"cocoyo/pkg/e"
	"cocoyo/pkg/setting"
	"cocoyo/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	gin.SetMode(setting.App.Key("RUN_MODE").String())

	router := routes.InitRouter()

	httpPort 		:=  ":" + setting.LoadServer().Key("HTTP_PORT").String()
	readTimeout 	:= setting.Server.Key("READ_TIMEOUT").MustInt(60)
	writeTimeout 	:= setting.Server.Key("WRITE_TIMEOUT").MustInt(60)

	s := &http.Server{
		Addr:               httpPort,
		Handler:        	router,
		ReadTimeout:    	time.Duration(readTimeout) * time.Second,
		WriteTimeout:   	time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: 	1 << 20,
	}

	e.New(e.INFO, fmt.Sprintf("start http server listening %s", httpPort))

	s.ListenAndServe()
}