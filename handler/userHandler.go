package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lfz757077613/goLearn/utils/myConf"
	"net/http"
	"sync"
)

type UserHandler struct {
}

var userHandlerOnce sync.Once
var userHandlerSingleton *UserHandler

func GetUserHandler() *UserHandler {
	userHandlerOnce.Do(func() {
		userHandlerSingleton = &UserHandler{}
	})
	return userHandlerSingleton
}

func (handler *UserHandler) HandleLogin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"result": myConf.GetInt("server","test",-1),
	})
}

func (handler *UserHandler) HandleIsLogin(context *gin.Context) {
}

func (handler *UserHandler) HandleRegister(context *gin.Context) {
}
