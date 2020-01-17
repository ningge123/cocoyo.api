package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func Setup()  {
	var err error
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "homestead", "secret", "192.168.10.10", "cocoyo"))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// 设置空闲中的最大连接数
	db.DB().SetMaxIdleConns(10)
	// 设置数据库的最大打开连接数。
	db.DB().SetMaxOpenConns(100)
	// 打开调试
	db.LogMode(true)
}

func CloseDB()  {
	defer db.Close()
}