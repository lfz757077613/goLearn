package midware

import (
	"github.com/lfz757077613/goLearn/utils/myLog"
	"io/ioutil"
	"net/http"
)

func MyParseParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			body, _ := ioutil.ReadAll(r.Body)
			myLog.Errorf("ParseForm error: [%s], url: [%s], body: [%s]", err, r.RequestURI, string(body))
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
