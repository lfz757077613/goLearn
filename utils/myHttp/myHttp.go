package myHttp

import (
	"crypto/tls"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: 3 * time.Second,
	// 返回301、302重定向时，不会自动发起重定向访问
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			// 不校验https证书
			InsecureSkipVerify: true,
		},
		MaxConnsPerHost:     300,
		MaxIdleConns:        150,
		MaxIdleConnsPerHost: 75,
		IdleConnTimeout:     10 * time.Second,
	},
}

func HttpPostWithRetry(url, body string, retryTime int) (result []byte, err error) {
	if retryTime < 0 {
		return nil, errors.New("HttpPostWithRetry retryTime is " + strconv.Itoa(retryTime))
	}
	for i := 0; i < retryTime; i++ {
		if result, err = HttpPost(url, body); err == nil {
			break
		}
	}
	return
}

func HttpGetWithRetry(url string, retryTime int) (result []byte, err error) {
	if retryTime < 0 {
		return nil, errors.New("HttpGetWithRetry retryTime is " + strconv.Itoa(retryTime))
	}
	for i := 0; i < retryTime; i++ {
		if result, err = HttpGet(url); err == nil {
			break
		}
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

func HttpGet(url string) ([]byte, error) {
	resp, err := httpClient.Get(url)
	// 重定向时，resp不为nil，如果CheckRedirect返回非ErrUseLastResponse错误，err也不为nil
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
