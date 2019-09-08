package main

import (
	"context"
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/lfz757077613/goLearn/handler"
	"github.com/lfz757077613/goLearn/utils/myConf"
	"github.com/lfz757077613/goLearn/utils/myLog"
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
	ginpprof.Wrap(r)
	r.Use(loggerMidware)
	// 拥有共同url前缀的的路由可以划为一个分组
	apiGroup := r.Group("/api")
	userHandler:= handler.UserHandler{}
	apiGroup.Any("/login", errWrapper(userHandler.HandleLogin))
	apiGroup.Any("/islogin", errWrapper(userHandler.HandleIsLogin))
	apiGroup.Any("/register", errWrapper(userHandler.HandleRegister))
	s := &http.Server{
		Addr:         ":" + myConf.GetString("server", "port", "8080"),
		Handler:      r,
		ReadTimeout:  time.Second * time.Duration(myConf.GetInt("server", "readTimeout", 3)),
		WriteTimeout: time.Second * time.Duration(myConf.GetInt("server", "writeTimeout", 3)),
	}
	go func() {
		myLog.Info("server start")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			myLog.Errorf("ListenAndServe error: [%s]", err)
			panic(err)
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
	myLog.Infof("quit signal received: [%s]", single.String())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		myLog.Errorf("Server Shutdown: [%s]", err)
	}
	myLog.Info("Server exit")
}

// 传入绑定了uid的log组件，打印请求整体日志
func loggerMidware(c *gin.Context) {
	start := time.Now()
	traceLog := myLog.WithField("uid", c.Query("uid"))
	c.Set("traceLog", traceLog)
	c.Next()
	traceLog.WithFields(logrus.Fields{
		"remoteIp": c.ClientIP(),
		"method":   c.Request.Method,
		"url":      c.Request.URL.Path + "?" + c.Request.URL.RawQuery,
		"status":   c.Writer.Status(),
		"cost":     fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
	}).Info("total log")
}

func errWrapper(handler func(c *gin.Context) error) func(*gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				myLog.GetUidTraceLog(c).Errorf("unknown panic: [%s]", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		err := handler(c)
		if err != nil {
			myLog.GetUidTraceLog(c).Errorf("unknown error: [%s]", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}
}