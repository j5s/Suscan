package models

import (
	"Suscan/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}


func Setup() {
	var err error
	db, err = gorm.Open(global.DatabaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.DatabaseSetting.UserName,
		global.DatabaseSetting.Password,
		global.DatabaseSetting.Host,
		global.DatabaseSetting.DBName))

	if err != nil {
		log.Fatalf("数据库配置错误: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return global.DatabaseSetting.TablePrefix + defaultTableName
	}
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(true)
}
