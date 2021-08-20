package models

import (
	"time"
)

type Setting struct {
	*Model
	Thread       string `json:"threadd"`
	Port        string `json:"port"`
	Cmd			string `json:"cmd"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

func GetSettingPort(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)
	return setting.Port
}

func GetSettingThread(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)

	return setting.Thread
}

func GetSettingCmd(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)
	return setting.Cmd
}

func EditSetting(data interface{}) bool {
	db.Model(&Setting{}).Updates(data)
	return true
}

func AddSetting(data map[string]interface{}) {
	setting := Setting{
		Thread:       data["thread"].(string),
		Port:        data["port"].(string),
		Cmd:		data["cmd"].(string),
		CreatedTime: time.Now().Format("20060102150405"),
		UpdatedTime: time.Now().Format("20060102150405"),
	}
	db.AutoMigrate(&setting)
	db.Create(&setting)
}

func EditSettingt(data interface{}) bool {
	GetSettingPort("port")
	db.Model(&Setting{}).Updates(data)
	return true
}