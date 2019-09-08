package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func (handler *UserHandler) HandleLogin(context *gin.Context) error {
	context.JSON(http.StatusOK,gin.H{
		"key":11,
	})
	return nil
}

func (handler *UserHandler) HandleIsLogin(context *gin.Context) error {
	return nil
}

func (handler *UserHandler) HandleRegister(context *gin.Context) error {
	return nil
}
