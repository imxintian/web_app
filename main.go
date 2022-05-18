package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/pkg/snowflake"
	"web_app/routes"
	"web_app/settings"
)

// go web 开发脚手架

func main() {
	// 1. 加载配置

	if len(os.Args) < 2 {
		fmt.Println("请输入配置文件名称")
		return
	}
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Println("加载配置失败", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Println("初始化数日志失败", err)
		return
	}
	defer zap.L().Sync()

	zap.L().Info("init logger success")

	// 3. 初始化数据库
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Println("初始化mysql失败", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("初始化redis失败", err)
		return
	}
	defer redis.Close()
	if err := snowflake.Init(settings.Conf.ServerConfig.StartTime, settings.Conf.ServerConfig.MachineID); err != nil {
		fmt.Println("初始化sf失败", err)
		return
	}
	// 4. 初始化路由
	r := routes.SetUp()

	// 5. 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.ServerConfig.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
