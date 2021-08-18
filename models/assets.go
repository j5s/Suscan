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

func GetAllAsset() (assets []Assets) {
	err := db.Select("target").Find(&assets)
	if err != nil {
		data := make(map[string]interface{})
		data["target"]="tuya.com"
		AddAsset(data)
		db.Select("target").Find(&assets)
	}
	return
}

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