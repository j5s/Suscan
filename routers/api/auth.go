package api

import (
	"Suscan/models"
	"Suscan/pkg/e"
	"Suscan/pkg/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := utils.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR
			} else {
				data["username"] = username
				data["token"] = token
				code = e.SUCCESS
			}

		} else {
			code = e.ERROR
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
