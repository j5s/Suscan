package v1

import (
	"Suscan/models"
	"Suscan/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ScanSetup(c *gin.Context) {
	thread := c.PostForm("thread")
	port := c.PostForm("port")
	cmd := c.PostForm("cmd")
	noping := c.PostForm("noping")
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["id"]=1
	data["thread"]=thread
	data["port"]=port
	data["cmd"]=cmd
	data["noping"]=noping

	nowTime := time.Now().Format("20060102150405")
	data["updated_time"] = nowTime
	models.EditSetting(data)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}