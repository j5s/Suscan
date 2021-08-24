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
	ResCode     string `json:"res_code"`
	ResResult   string `json:"res_result"`
	ResType     string `json:"res_type"`
	ResUrl      string `json:"res_url"`
	ResTitle    string `json:"res_title"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

func ExistIplist(ip, port string) (bool, int) {
	var iplist Iplist
	db.Select("id").Where("ip = ? and port = ? ", ip, port).First(&iplist)

	if iplist.ID > 0 {
		return true, iplist.ID
	}
	return false, iplist.ID
}

func EditIplist(id int, data interface{}) bool {
	db.Model(&Iplist{}).Where("id = ?", id).Updates(data)
	return true
}

func AddIplist(data map[string]interface{}) {
	nowTime := time.Now().Format("20060102150405")
	IpList := Iplist{
		Url:         data["url"].(string),
		Ip:          data["ip"].(string),
		Port:        data["port"].(string),
		State:       data["state"].(string),
		Protocol:    data["protocol"].(string),
		Service:     data["service"].(string),
		ResCode:     data["res_code"].(string),
		ResResult:   data["res_result"].(string),
		ResType:     data["res_type"].(string),
		ResUrl:      data["res_url"].(string),
		ResTitle:    data["res_title"].(string),
		CreatedTime: nowTime,
		UpdatedTime: nowTime,
	}
	db.AutoMigrate(&IpList)
	db.Create(&IpList)
}
