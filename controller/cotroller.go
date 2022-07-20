package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {

}

//新增用户

func (controller *UserController) UserAdd(context *gin.Context)  {

	name, exist := context.GetPostForm("name")
	if !exist || name == ""{
		context.JSON(http.StatusOK, gin.H{
			"name": "请输入你的名字：name",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

//query
func (controller *UserController) UserQuery(context *gin.Context)  {
	id := context.Query("id")
	println("id: ", id)
	if id >= "100" {
		context.JSON(403, gin.H{
			"msg": "id错误",
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}