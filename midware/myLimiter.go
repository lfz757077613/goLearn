package midware

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/lfz757077613/goLearn/utils/myConf"
//	"github.com/lfz757077613/goLearn/utils/myLog"
//	"net/http"
//)
//
//func MyLimiter(c *gin.Context) {
//	url := c.Request.URL.Path
//	if limiter := myConf.GetLimiterByUrl(url); limiter != nil && !limiter.Allow() {
//		myLog.GetUidTraceLog(c).Errorf("limit url: [%s]", url)
//		c.AbortWithStatus(http.StatusForbidden)
//		return
//	}
//	c.Next()
//}
