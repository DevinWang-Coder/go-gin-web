package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-web/controller"
)

func RegisterRouter(routers *gin.Engine)  {
	routerUser(routers)
}

//用户路由
func routerUser(engine *gin.Engine)  {
	var group = engine.Group("/api/user")
	{
		con := &controller.UserController{}
		group.POST("/userAdd", con.UserAdd)
		group.GET("/userQuery", con.UserQuery)
	}
}