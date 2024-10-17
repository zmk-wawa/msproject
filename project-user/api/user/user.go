package user

import "github.com/gin-gonic/gin"

// 为了将 ctx *gin.Context 方法抽出来
type HandlerUser struct {
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	ctx.JSON(200, "getCaptcha success")
}
