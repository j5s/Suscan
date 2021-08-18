package routers

import (
	"Suscan/routers/api"
	v1 "Suscan/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/v1/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")

	//端口扫描 Nmap
	//http://localhost:18000/api/v1/scan
	apiv1.GET("/scan", v1.ScanPort)

	//添加扫描资产
	//http://localhost:18000/api/v1/assets
	apiv1.POST("/assets", v1.AssetsSetup)

	//查找所有要扫描的资产
	//http://localhost:18000/api/v1/getassets
	//apiv1.GET("/getassets", v1.GetAssets)



	return r
}