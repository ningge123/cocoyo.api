package main

import (
	"cocoyo/models"
	"cocoyo/pkg/e"
	"cocoyo/pkg/setting"
	"cocoyo/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init()  {
	models.Setup()
}

func main() {
	rubMode := setting.Cfg.Section("").Key("RUN_MODE").String()
	gin.SetMode(rubMode)

	router := routers.InitRouter()

	httpPort 		:= ":" + setting.Cfg.Section("server").Key("HTTP_PORT").String()
	readTimeout 	:= setting.Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)
	writeTimeout 	:= setting.Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60)

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