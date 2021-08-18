package v1

import (
	"Suscan/models"
	"Suscan/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)


func AssetsSetup(c *gin.Context) {
	assets := c.PostForm("assets")
	code := e.SUCCESS
	fmt.Println(assets)
	assetResult := strings.Split(assets,"\n")

	for _,as := range assetResult{
		data := make(map[string]interface{})
		data["target"]=as
		//资产对比
		ok, id := models.ExistAsset(as)
		if ok {
			dataUpdate := make(map[string]interface{})
			nowTime := time.Now().Format("20060102150405")
			dataUpdate["updated_time"] = nowTime
			models.EditAsset(id, dataUpdate)
		} else {
			models.AddAsset(data)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func GetAssets(c *gin.Context) {
	code := e.SUCCESS
	result := models.GetAllAsset()

	for _,r := range result{
		fmt.Println(r.Target)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

