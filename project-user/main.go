package main

import (
	"github.com/gin-gonic/gin"
	srv "project-common"
	_ "test.com/project-user/api"
	"test.com/project-user/config"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	// 初始化日志
	// 从配置中读取日志配置，初始化日志
	config.C.InitZapLog()

	// 先注册路由，再Run进行优雅启停
	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr)

	// 将下面的直接调用包装几层
	//r.POST("project/login/getCaptcha", func(ctx *gin.Context) {
	//	ctx.JSON(200, "getCaptcha success")
	//})

}
