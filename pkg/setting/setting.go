package setting

import (
	"cocoyo/pkg/e"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init() {
	var err error

	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		e.New(e.ERROR, fmt.Sprintf("Fail to parse 'conf/app.ini': %v", err))
	}

	//LoadBase()
	//LoadServer()
	//LoadApp()
}

func LoadBase()  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer()  {
	cfg, err := Cfg.GetSection("server")

	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = cfg.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration( cfg.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration( cfg.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp()  {
	cfg, err := Cfg.GetSection("app")

	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret =  cfg.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize =  cfg.Key("PAGE_SIZE").MustInt(10)
}
