package handler

import (
	"github.com/lfz757077613/goLearn/utils/myLog"
	"io"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if _, err := io.WriteString(w, "ok"); err != nil {
		myLog.Errorf("handleCheckPreload error: [%s]", err)
	}
}
