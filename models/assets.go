package models

import "time"

type Assets struct {
	*Model
	Target      string `json:"target"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}


func AddAsset(data map[string]interface{}) {
	assets := Assets{
		Target:      data["target"].(string),
		CreatedTime: time.Now().Format("20060102150405"),
		UpdatedTime: time.Now().Format("20060102150405"),
	}
	db.AutoMigrate(&assets)
	db.Create(&assets)
}

func ExistAsset(target string) (bool, int) {
	var assets Assets
	db.Select("id").Where("target = ? ", target).First(&assets)
	if assets.ID > 0 {
		return true, assets.ID
	}
	return false, assets.ID
}

func EditAsset(id int, data interface{}) bool {
	db.Model(&Assets{}).Where("id = ?", id).Updates(data)
	return true
}

//查询所有target

func GetAllAsset() (assets []Assets) {
	db.Select("target").Find(&assets)
	return
}

//返回自定义的高危端口

func GetPortResult() (iplist []Iplist) {
	dbTmp := db
	dbTmp.Where("port = ? OR port = ? OR port = ? OR port = ? OR port = ? OR port = ? OR port = ? OR port = ? OR port = ? OR port = ?", 3389, 22, 1988,21, 3306, 6379, 5200, 446, 7799, 33033).Find(&iplist)
	return
}

//返回自定义的高危协议

func GetProResult() (iplist []Iplist) {
	dbTmp := db
	dbTmp.Where("service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ? OR service = ?", "mysql", "mssql", "redis", "memcache", "mongo", "ftp", "tftp", "ssh", "weblogic", "websphere", "tomcat", "ms-wbt-server", "oracle", "afs3-callback").Find(&iplist)
	return
}


