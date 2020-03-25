package setting

import (
	"cocoyo/pkg/e"
	"fmt"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
	App *ini.Section
	Server *ini.Section
	Database *ini.Section
	Filesystem *ini.Section
	Jwt *ini.Section
)

func init() {
	var err error

	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		e.New(e.ERROR, fmt.Sprintf("Fail to parse 'conf/app.ini': %v", err))
	}

	App 		= LoadApp()
	Server 		= LoadServer()
	Database 	= LoadDatabase()
	Filesystem 	= LoadFilesystem()
	Jwt 		= LoadJwt()
}

func LoadApp() *ini.Section {
	return Cfg.Section("app")
}

func LoadServer() *ini.Section {
	return Cfg.Section("server")
}

func LoadDatabase() *ini.Section {
	return Cfg.Section("database")
}

func LoadFilesystem() *ini.Section {
	return Cfg.Section("filesystems")
}

func LoadJwt() *ini.Section {
	return Cfg.Section("jwt")
}