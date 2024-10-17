package router

import (
	"github.com/gin-gonic/gin"
	"test.com/project-user/api/user"
)

//  类似Service接口层
// 为了将r.POST路由抽象成公共代码，需要将其变成接口

type Router interface {
	Route(r *gin.Engine)
}

//注册接口的类，不同的路由请求都要调用其来注册
type RegisterRouter struct {
}

// 构造函数
func NewRegisterRouter() *RegisterRouter {
	return &RegisterRouter{}
}

// 用调用接口方法
func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

// 在此注册路由
func InitRouter(r *gin.Engine) {
	router := NewRegisterRouter()
	router.Route(&user.RouterUser{}, r)
}
