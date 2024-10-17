package user

import "github.com/gin-gonic/gin"

// 实现Router接口中的方法

// 实现接口方法
// 定义类，然后依托类来实现方法
type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.POST("project/login/getCaptcha", h.getCaptcha)
}
