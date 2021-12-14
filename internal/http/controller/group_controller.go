package controller

import (
	"github.com/gin-gonic/gin"
)

type GroupController struct {
	BaseController
}

func (c *GroupController) Create(content *gin.Context) {
	c.JsonSuccess(content, nil, "")
}

func (c *GroupController) Delete(content *gin.Context) {
	c.JsonSuccess(content, nil, "")
}

func (c *GroupController) Info(content *gin.Context) {
	c.JsonSuccess(content, nil, "")
}

func (c *GroupController) All(content *gin.Context) {
	c.JsonSuccess(content, nil, "")
}

func (c *GroupController) Join(content *gin.Context) {
	c.JsonSuccess(content, nil, "")
}

func (c *GroupController) Quit(content *gin.Context) {
	c.JsonSuccess(content, nil, "")
}

func (c *GroupController) List(content *gin.Context) {
	c.JsonSuccess(content, nil, "")
}
