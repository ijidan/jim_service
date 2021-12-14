package controller

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Pong(content *gin.Context) {
	c.JsonSuccess(content, gin.H{"message": "pong"}, "")
}
