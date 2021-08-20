package v1

import (
	"Suscan/models"
	"Suscan/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取高危端口IP的相关信息
func Getvulport(c *gin.Context) {
	data := make(map[string]interface{})
	data["ip"] = models.GetPortResult()
	fmt.Println(data)
	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取高危协议IP的相关信息
func Getvulpro(c *gin.Context) {
	data := make(map[string]interface{})
	data["pro"] = models.GetProResult()
	fmt.Println(data)
	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
