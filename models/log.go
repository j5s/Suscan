package models

import "time"

type Log struct {
	Id          int    `gorm:"primary_key" json:"id"`
	TaskName    string `json:"task_name"`
	TaskType    string `json:"task_type"`
	AllNum      int    `json:"all_num"`
	RunTime     string `json:"run_time"`
	CreatedTime string `json:"created_time"`
}



func AddLog(data map[string]interface{})int {
	log := Log{
		TaskName:    data["task_name"].(string),
		TaskType:    data["task_type"].(string),
		AllNum:      data["all_num"].(int),
		RunTime:     data["run_time"].(string),
		CreatedTime: time.Now().Format("20060102150405"),
	}
	db.AutoMigrate(&log)
	db.Create(&log)
	return log.Id
}
