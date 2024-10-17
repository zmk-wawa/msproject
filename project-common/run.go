package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, srvName string, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 保证优雅启停
	// 定义了协程，不会阻塞，所以需要手动阻塞，不然就结束了
	go func() {
		log.Printf("%s running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	// 创建os.Signal类型通道，用于接受信号
	quit := make(chan os.Signal)
	// 识别 SIGINT(CTRL+C) 与 SIGTERM(结束程序)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//接受操作符，阻塞程序，从quit通道中接受数据，直到收到上面的两种信号
	<-quit
	// 能到说明收到了上面两种信号之一，即结束程序
	log.Printf("Shutting Down project %s....\n", srvName)
	// 创建一个带有2s超市的上下文ctx,太短会导致超时操作无法执行完
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// 在main函数返回时，调用cancel()函数，来确保srv.Shutdown完成或超时后取消上下文
	// defer会延迟函数执行直至包含defer语句的函数执行完毕
	defer cancel()

	// tips: 这里可以放其他需要关闭的，比如grpc等服务，如果需要

	// 使用Shutdown方法，使用上下文ctx优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		// 关闭时出现错误则打印
		log.Fatalf("%s Shutdown, cause by : %v\n", srvName, err)
	}

	select {
	// 检查上下文的Done通道，如果超时，则记录超时消息
	case <-ctx.Done():
		log.Println("wait timeout...")
	}
	// 到这里说明成功关闭，打印后会调用cancel()函数，取消上下文
	log.Printf("%s stop success...\n", srvName)
}
