package routers

import (
	"Suscan/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/v1/auth", api.GetAuth)
	//apiv1 := r.Group("/api/v1")

	return r
}