package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/lemoba/go-admin/model"
)

func GetUsers(c *gin.Context) {
	data := model.GetUsers()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func AddUser(c *gin.Context) {

}

func ShowUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DestoryUser(c *gin.Context) {

}
