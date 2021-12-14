package controller

import (
	"github.com/gin-gonic/gin"
	"jim/global"
	"net/http"
)

type BaseController struct {
}

func (c *BaseController) JsonSuccess(context *gin.Context, data map[string]interface{}, jumpUrl string) {
	rsp := global.Response.JsonSuccess(data, jumpUrl)
	context.JSON(http.StatusOK, rsp)
	return
}

func (c *BaseController) JsonFail(context *gin.Context, code int32, businessCode int32, message string, data map[string]interface{}, jumpUrl string) {
	rsp := global.Response.JsonFail(code, businessCode, message, data, jumpUrl)
	context.JSON(http.StatusOK, rsp)
	return
}
