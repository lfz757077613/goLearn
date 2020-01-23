package midware

import (
	"github.com/lfz757077613/goLearn/utils/myLog"
	"net/http"
	"runtime/debug"
)

func MyRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				myLog.Errorf("unknown panic: [%s], stacktrace: [%s]", err, debug.Stack())
				http.Error(w, "", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
