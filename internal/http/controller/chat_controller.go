package controller

import (
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	BaseController
}

func (c *ChatController) History(content *gin.Context) {
	c.JsonSuccess(content, nil, "")
}
