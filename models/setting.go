package models

type Setting struct {
	*Model
	Thread       string `json:"thread"`
	Port         string `json:"port"`
	Cmd          string `json:"cmd"`
	NoPing          string `json:"noping"`
	CreatedTime  string `json:"created_time"`
	UpdatedTime  string `json:"updated_time"`
}


func GetSettingPort(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)
	return setting.Port
}

func GetSettingCmd(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)
	return setting.Cmd
}

func GetSettingNoPing(maps interface{}) (a string) {
	var setting Setting
	db.Where(maps).First(&setting)
	return setting.NoPing
}

func EditSetting(data interface{}) bool {
	db.Model(&Setting{}).Updates(data)
	return true
}