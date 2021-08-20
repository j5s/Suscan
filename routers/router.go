package routers

import (
	v1 "Suscan/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//登陆认证，写前端的时候再加
	//r.GET("/api/v1/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")


	//添加扫描资产
	//http://localhost:18000/api/v1/assets
	apiv1.POST("/assets", v1.AssetsSetup)

	//端口扫描 Nmap
	//http://localhost:18000/api/v1/scan
	apiv1.GET("/scan", v1.ScanPort)

	//修改扫描设置
	//http://localhost:18000/api/v1/scansetting
	apiv1.POST("/scansetting", v1.ScanSetup)

	//获取高危端口
	//http://localhost:18000/api/v1/getVulPort
	apiv1.GET("/getVulPort", v1.Getvulport)

	//获取高危协议
	//http://localhost:18000/api/v1/getVulPort
	apiv1.GET("/getVulPro", v1.Getvulpro)

	return r
}