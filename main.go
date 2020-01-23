package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/lfz757077613/goLearn/handler"
	"github.com/lfz757077613/goLearn/midware"
	"github.com/lfz757077613/goLearn/utils/myConf"
	"github.com/lfz757077613/goLearn/utils/myLog"
	"github.com/lfz757077613/goLearn/utils/shutDownhook"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.Use(midware.MyRecover, midware.MyLogger, midware.MyParseParam)
	r.HandleFunc("/next", handler.HandleLogin)
	s := &http.Server{
		Addr:         ":" + myConf.GetString("server", "port", "8080"),
		Handler:      http.TimeoutHandler(r, 3*time.Second, ""),
		ReadTimeout:  time.Second * time.Duration(myConf.GetInt("server", "readTimeout", 3)),
		WriteTimeout: time.Second * time.Duration(myConf.GetInt("server", "writeTimeout", 3)),
	}
	go func() {
		myLog.Info("server start")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			myLog.Panicf("ListenAndServe error: [%s]", err)
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
		myLog.Errorf("Server Shutdown error: [%s]", err)
	}
	shutDownhook.RunShutdownHook()
	myLog.Info("Server exit")
}
