package main

import (
	"github.com/gin-gonic/gin"
	srv "project-common"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	// 先注册路由，再Run进行优雅启停
	router.InitRouter(r)
	srv.Run(r, "project-user", ":80")

	// 将下面的直接调用包装几层
	//r.POST("project/login/getCaptcha", func(ctx *gin.Context) {
	//	ctx.JSON(200, "getCaptcha success")
	//})

}
