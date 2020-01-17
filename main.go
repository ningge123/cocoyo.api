package main

import (
	"cocoyo/models"
	"cocoyo/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init()  {
	models.Setup()
}

func main() {
	gin.SetMode("debug")

	router := routers.InitRouter()

	httpPort := ":8080"
	readTimeout := 60
	writeTimeout := 60

	s := &http.Server{
		Addr:               httpPort,
		Handler:        	router,
		ReadTimeout:    	time.Duration(readTimeout) * time.Second,
		WriteTimeout:   	time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening %s", httpPort)

	s.ListenAndServe()
}