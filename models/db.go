package models

import (
	"cocoyo/pkg/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init()  {
	var err error

	db, err = gorm.Open(
		setting.Database.Key("DB_CONNECTION").String(),
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.Database.Key("DB_USERNAME").String(),
			setting.Database.Key("DB_PASSWORD").String(),
			setting.Database.Key("DB_HOST").String() + ":" + setting.Database.Key("DB_PORT").String(),
			setting.Database.Key("DB_DATABASE").String()))

	if err != nil {
		panic(fmt.Sprintf("models.Setup err: %v", err))
	}

	// 设置空闲中的最大连接数
	db.DB().SetMaxIdleConns(10)
	// 设置数据库的最大打开连接数。
	db.DB().SetMaxOpenConns(100)
	// 打开调试
	db.LogMode(true)
}

func ReturnDB() *gorm.DB {
	return db
}

