package main

import (
	"context"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/lfz757077613/goLearn/handler"
	"github.com/lfz757077613/goLearn/midware"
	"github.com/lfz757077613/goLearn/utils/myConf"
	"github.com/lfz757077613/goLearn/utils/shutDownhook"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	r := gin.New()
	r.RedirectTrailingSlash = false
	pprof.Register(r)
	r.Use(midware.MyLogger, midware.MyRecover)
	// 拥有共同url前缀的的路由可以划为一个分组
	apiGroup := r.Group("/api")
	userHandler := handler.UserHandler{}
	apiGroup.Any("/login", userHandler.HandleLogin)
	apiGroup.Any("/islogin", userHandler.HandleIsLogin)
	apiGroup.Any("/register", userHandler.HandleRegister)
	s := &http.Server{
		Addr:         ":" + myConf.GetString("server", "port", "8080"),
		Handler:      r,
		ReadTimeout:  time.Second * time.Duration(myConf.GetInt("server", "readTimeout", 3)),
		WriteTimeout: time.Second * time.Duration(myConf.GetInt("server", "writeTimeout", 3)),
	}
	go func() {
		logrus.Info("server start")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Panicf("ListenAndServe error: [%s]", err)
		}
	}()
	// 优雅关机
	waitShutDownSignal(s)
}

func waitShutDownSignal(s *http.Server) {
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// 要用缓冲channel，否则可能丢失信号
	quitSignalChan := make(chan os.Signal, 1)
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quitSignalChan, syscall.SIGINT, syscall.SIGTERM)
	single := <-quitSignalChan
	logrus.Infof("quit signal received: [%s]", single.String())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logrus.Errorf("Server Shutdown error: [%s]", err)
	}
	shutDownhook.RunShutdownHook()
	logrus.Info("Server exit")
}
