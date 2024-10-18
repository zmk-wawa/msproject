package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-user/router"
)

func init() {
	// 将这里(user)要注册的路由批量注册进去
	log.Println("init user router")
	// 传入一个或多个实现了Router接口的类
	router.Register(&RouterUser{})
}

// 实现Router接口中的方法

// 实现接口方法
// 定义类，然后依托类来实现方法
type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := NewHandlerUser() // 调用自己写的构造函数
	r.POST("project/login/getCaptcha", h.getCaptcha)
}
