package router

import (
	"github.com/gin-gonic/gin"
)

//  类似Service接口层
// 为了将r.POST路由抽象成公共代码，需要将其变成接口

type Router interface {
	Route(r *gin.Engine)
}

////注册接口的类，不同的路由请求都要调用其来注册
//type RegisterRouter struct {
//}
//
//// 构造函数
//func NewRegisterRouter() *RegisterRouter {
//	return &RegisterRouter{}
//}
//
//// 用调用接口方法
//func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
//	ro.Route(r)
//}

// 用来批量注册路由
var routers []Router

// 在此注册路由
func InitRouter(r *gin.Engine) {
	//router := NewRegisterRouter()
	//router.Route(&user.RouterUser{}, r)
	// 批量注册
	for _, ro := range routers {
		ro.Route(r)
	}
}

// 上层调用，将传入的Routers加进routers，等待最后批量注册
func Register(ro ...Router) {
	routers = append(routers, ro...)
}
