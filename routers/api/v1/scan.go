package v1

import (
	"Suscan/pkg/e"
	"Suscan/pkg/nmap"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 开始nmap端口扫描
func ScanPort(c *gin.Context) {
	code := e.SUCCESS
	nmap.NmapScan()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}