package user

import (
	"github.com/gin-gonic/gin"
	srv "project-common"
)

// 为了将 ctx *gin.Context 方法抽出来
type HandlerUser struct {
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	// 调用common模块中的标准模板来生成要输出返回JSON
	res := &srv.Result{}
	ctx.JSON(200, res.Success("test success"))
}
