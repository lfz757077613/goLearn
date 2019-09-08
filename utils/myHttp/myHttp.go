package myHttp

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: 3 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}


func HttpPostWithRetry(url, body string, retryTime int) (result []byte, err error) {
	for i := 0; i < retryTime; i++ {
		result, err = HttpPost(url, body)
	}
	return
}

func HttpPost(url, body string) ([]byte, error) {
	resp, err := httpClient.Post(url, gin.MIMEPOSTForm, strings.NewReader(body))
	// 重定向时，resp不为nil，如果CheckRedirect返回非ErrUseLastResponse错误，err也不为nil
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

