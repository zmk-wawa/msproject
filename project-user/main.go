package main

import (
	"github.com/gin-gonic/gin"
	"log"
	srv "project-common"
	"project-common/logs"
	_ "test.com/project-user/api"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	// 初始化日志
	// 从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "/Users/zmk/Downloads/go/project/ms_project/project-common/logs/project-debug.log",
		InfoFileName:  "/Users/zmk/Downloads/go/project/ms_project/project-common/logs/project-info.log",
		WarnFileName:  "/Users/zmk/Downloads/go/project/ms_project/project-common/logs/project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}

	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}

	// 先注册路由，再Run进行优雅启停
	router.InitRouter(r)
	srv.Run(r, "project-user", ":80")

	// 将下面的直接调用包装几层
	//r.POST("project/login/getCaptcha", func(ctx *gin.Context) {
	//	ctx.JSON(200, "getCaptcha success")
	//})

}
