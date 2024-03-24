package main

import (
	"context"
	"fmt"
	"gin_demo2/controllers"
	"gin_demo2/dao/mysql"
	"gin_demo2/dao/redis"
	"gin_demo2/logger"
	"gin_demo2/pkg/snowflake"
	"gin_demo2/routes"
	"gin_demo2/settings"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// go web开发较为通用的脚手架模板
func main() {
	//1.加载配置，viper库中结构体进行序列化和反序列化，都使用的标签是mapstructure，而不是yaml，json啥的
	if err := settings.Init(); err != nil {
		fmt.Println("配置文件读取失败：", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Println("日志文件设置失败：", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3.初始化mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Println("mysql设置失败：", err)
		return
	}
	defer mysql.Close()
	//4.初始化redis连接
	if err := redis.Init(); err != nil {
		fmt.Println("redis设置失败：", err)
		return
	}
	defer redis.Close()

	//使用雪花算法生成用户id
	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Println("id生成失败：", err)
		return
	}
	//初始化gin框架内置校验器里的翻译器，一般来说err都是英文，有了翻译器可以变为中文
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err:%v\n", err)
		return
	}
	//5.注册路由
	r := routes.SetupRouter()
	s := fmt.Sprintf("%s%d", ":", viper.GetInt("app.port")) //port 8080
	if err := r.Run(s); err != nil {
		fmt.Printf("run server failed,err:%v\n", err)
		return
	}
	//6.启动服务(优雅开关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
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
