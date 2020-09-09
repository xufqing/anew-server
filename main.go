package main

import (
	"anew-server/common"
	"anew-server/initialize"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.Logger()
	// 初始化路由
	r := initialize.Routers()

	// 启动服务器
	host := "0.0.0.0"
	port := common.Conf.System.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}
	go func() {
		// 加入pprof性能分析
		if err := http.ListenAndServe(":8005", nil); err != nil {
			common.Log.Error("listen pprof error: ", err)
		}
	}()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			common.Log.Error("listen error: ", err)
		}
	}()

	common.Log.Info(fmt.Sprintf("Server is running at %s:%d", host, port))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	common.Log.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		common.Log.Error("Server forced to shutdown: ", err)
	}

	common.Log.Info("Server exiting")
}