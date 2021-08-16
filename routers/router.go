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

	//http://localhost:18000/api/v1/scan
	apiv1.GET("/scan", v1.ScanPort)

	return r
}