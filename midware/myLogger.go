package midware

import (
	"fmt"
	"github.com/lfz757077613/goLearn/utils/myLog"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func MyLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		clientIP, _, _ := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
		next.ServeHTTP(w, r)
		// w是*http.timeoutWriter，当使用了http.TimeoutHandler时
		statusField := reflect.ValueOf(w).Elem().FieldByName("code")
		var status int64
		if !statusField.IsValid() {
			// w是*http.response
			statusField = reflect.ValueOf(w).Elem().FieldByName("status")
		}
		if statusField.IsValid() {
			status = statusField.Int()
		}
		body, _ := ioutil.ReadAll(r.Body)
		myLog.WithFields(logrus.Fields{
			"remoteIp": clientIP,
			"method":   r.Method,
			"url":      r.RequestURI,
			"bode":     string(body),
			"status":   status,
			"cost":     fmt.Sprintf("%dms", time.Since(start).Nanoseconds()/1e6),
		}).Info("total log")
	})
}

