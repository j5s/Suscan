package models

import (
	"time"
)

type Iplist struct {
	*Model
	Url         string `json:"url"`
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	State       string `json:"state"`
	Protocol    string `json:"protocol"`
	Service     string `json:"service"`
	Res_code    string `json:"res_code"`
	Res_result  string `json:"res_result"`
	Res_type    string `json:"res_type"`
	Res_url     string `json:"res_url"`
	Res_title   string `json:"res_title"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

func ExistIplist(ip, port string) (bool, int) {
	var iplist Iplist
	db.Select("id").Where("ip = ? and port = ? ", ip, port).First(&iplist)
	//如果返回的id>0，也就是数据库里存在过了数据
	if iplist.ID > 0 {
		return true, iplist.ID
	}
	return false, iplist.ID
}

func EditIplist(id int, data interface{}) bool {
	db.Model(&Iplist{}).Where("id = ?", id).Updates(data)
	return true
}

//创建任务，返回任务id
func AddIplist(data map[string]interface{}) {
	nowTime := time.Now().Format("20060102150405")
	iplist := Iplist{
		Url:         data["url"].(string),
		Ip:          data["ip"].(string),
		Port:        data["port"].(string),
		State:       data["state"].(string),
		Protocol:    data["protocol"].(string),
		Service:     data["service"].(string),
		Res_code:    data["res_code"].(string),
		Res_result:  data["res_result"].(string),
		Res_type:    data["res_type"].(string),
		Res_url:     data["res_url"].(string),
		Res_title:   data["res_title"].(string),
		CreatedTime: nowTime,
		UpdatedTime: nowTime,
	}
	db.AutoMigrate(&iplist)
	db.Create(&iplist)
}
