package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	srv "project-common"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
	"time"
)

// 为了将 ctx *gin.Context 方法抽出来
type HandlerUser struct {
	// 要缓存验证码，故在这里接入cache
	cache repo.Cache
}

// 为了方便缓存的管理，新增HandleUser的构造函数
// 这样后续需要调整缓存所使用的数据库，就在构造函数中更换
func NewHandlerUser() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	// 调用common模块中的标准模板来生成要输出返回JSON
	res := &srv.Result{}
	// 处理前端在注册时，点击生成验证码后：返回手机号验证码的逻辑
	// 1. 获取手机号，并校验手机号的合法性
	mobileNum := ctx.PostForm("mobile")
	if !srv.VerifyMobile(mobileNum) {
		// 不合法
		log.Println("手机号不合法")
		ctx.JSON(http.StatusOK, res.Fail(model.NoLegalMobile, "mobile number is illegal"))
		return
	}
	// 2. 生成验证码（4位 or 6位随机数）
	code := "123456"
	// 3. 调用短信平台（三方 放入go协程中，不阻塞主协程+接口可以快速响应+并发处理）
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("短信平台调用成功，发送短信")
		// 4. 将手机号与验证码<k,v>存入Redis中，设定过期时间（eg:15min）
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := h.cache.Put(c, "Register_"+mobileNum, code, 15*time.Minute); err != nil {
			// 存储失败
			log.Printf("存入缓存失败, 原因是 %v \n", err)
		}
		log.Printf("成功将手机号和验证码存入Redis数据库: Register_%s : %s", mobileNum, code)
	}()
	ctx.JSON(http.StatusOK, res.Success(code))
}
