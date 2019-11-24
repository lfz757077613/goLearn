package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// 传入绑定了uid的log组件，打印请求整体日志
func MyLogger(c *gin.Context) {
	start := time.Now()
	traceLog := logrus.WithField("uid", c.Query("uid"))
	c.Set("traceLog", traceLog)
	c.Next()
	rawQuery := ""
	if c.Request.URL.RawQuery != "" {
		rawQuery = "?" + c.Request.URL.RawQuery
	}
	traceLog.WithFields(logrus.Fields{
		"remoteIp": c.ClientIP(),
		"method":   c.Request.Method,
		"url":      c.Request.URL.Path + rawQuery,
		"status":   c.Writer.Status(),
		"cost":     fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
	}).Info("total log")
}
