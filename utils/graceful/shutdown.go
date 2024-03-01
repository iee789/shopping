package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 合理的停止server

// 参数含义：1.需要关闭的HTTP server实例；2.关闭server的超时时间
func ShutdownGin(instance *http.Server, timeout time.Duration) {
	// 当收到一个 os.Interrupt 或者 syscall.SIGTERM 信号.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 监听系统信号
	<-quit
	// 1.打印日志关闭服务器
	log.Println("关闭 Server ...")
	// 2.创建上下文设置超时时间保证server关闭操作在指定时间完成
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 3.调用关闭函数
	if err := instance.Shutdown(ctx); err != nil {
		// 4.如果关闭失败则打印日志
		log.Fatal("Server 关闭:", err)
	}
	// 超时5秒 ctx.Done().
	select {
		// 5.指定时间内如果服务器没完全退出则打印日志提示超时
	case <-ctx.Done():
		log.Println("超时5秒.")
	}
	// 6.最后打印日志提示退出
	log.Println("Server 退出")
}
