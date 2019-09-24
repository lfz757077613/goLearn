package handler

import (
	"github.com/gin-gonic/gin"
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
		"key": 11,
	})
}

func (handler *UserHandler) HandleIsLogin(context *gin.Context) {
}

func (handler *UserHandler) HandleRegister(context *gin.Context) {
}
