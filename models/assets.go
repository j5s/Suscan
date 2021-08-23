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

//增加setting表内容
//data := make(map[string]interface{})
//data["thread"] = "2000"
////第一次没有配置就默认配置常用端口
//data["port"] = "80,1433,1521,1583,2100,2049,3050,3306,3351,5000,5432,5433,5601,5984,6082,6379,7474,8080,8088,8089,8098,8471,9000,9160,9200,9300,9471,11211,15672,19888,27017,27019,27080,28017,50000,50070,50090"
////最多同时启动5个nmap扫描终端
//data["cmd"] = "5"
//AddSetting(data)
//db.Where(maps).First(&setting)


func ExistAsset(target string) (bool, int) {
	var assets Assets
	db.Select("id").Where("target = ? ", target).First(&assets)
	//如果返回的id>0，也就是数据库里存在过了数据
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

//提供给安全平台查询的接口
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


