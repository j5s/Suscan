package models

type Setting struct {
	*Model
	Thread       string `json:"thread"`
	Port         string `json:"port"`
	Cmd          string `json:"cmd"`
	Timetemplate string `json:"timetemplate"`
	CreatedTime  string `json:"created_time"`
	UpdatedTime  string `json:"updated_time"`
}

func GetSettingPort(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)
	return setting.Port
}

func GetSettingTiming(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)
	return setting.Timetemplate
}

func GetSettingCmd(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)
	return setting.Cmd
}

func EditSettingt(data interface{}) bool {
	db.Model(&Setting{}).Updates(data)
	return true
}

//func AddSetting(data map[string]interface{}) {
//	setting := Setting{
//		Thread:      data["thread"].(string),
//		Port:        data["port"].(string),
//		Cmd:         data["cmd"].(string),
//		Timetemplate:  data["timetemplate"].(string),
//		CreatedTime: time.Now().Format("20060102150405"),
//		UpdatedTime: time.Now().Format("20060102150405"),
//	}
//	db.AutoMigrate(&setting)
//	db.Create(&setting)
//}

////增加setting表内容
//data := make(map[string]interface{})
//data["thread"] = "2000"
////第一次没有配置就默认配置常用端口
//data["port"] = "80,1433,1521,1583,2100,2049,3050,3306,3351,5000,5432,5433,5601,5984,6082,6379,7474,8080,8088,8089,8098,8471,9000,9160,9200,9300,9471,11211,15672,19888,27017,27019,27080,28017,50000,50070,50090"
////最多同时启动5个nmap扫描终端
//data["cmd"] = "5"
//data["timetemplate"] = 5
//AddSetting(data)