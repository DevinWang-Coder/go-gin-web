package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-web/routers"
)

func main()  {
	//create
	r := gin.Default()

	//添加路由
	routers.RegisterRouter(r)
	r.Run(":8888")

}