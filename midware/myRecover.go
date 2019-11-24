package midware

import (
	"github.com/gin-gonic/gin"
	"github.com/lfz757077613/goLearn/utils/myLog"
	"net/http"
	"runtime/debug"
)

func MyRecover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			myLog.GetUidTraceLog(c).Errorf("unknown panic: [%s], stacktrace: [%s]", err, debug.Stack())
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}