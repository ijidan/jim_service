package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/util/grand"
	"jim/global"
	"jim/internal/http/request"
	"jim/internal/model"
	"net/http"
)

type UserController struct {
	BaseController
}

func (c *UserController) Register(content *gin.Context) {
	var req request.UserRegisterRequest
	if err := content.ShouldBindJSON(&req); err != nil {
		c.JsonFail(content, http.StatusBadRequest, 0, err.Error(), nil, "")
	}

	var user model.User
	global.Db.Where("nickname=?", req.Nickname).First(&user)
	if user.ID != 0 {
		c.JsonFail(content, http.StatusNotFound, 0, "用户已存在", nil, "")
	}
	key := grand.Letters(10)
	encryptedPassword, err := gaes.Encrypt([]byte(req.Password), []byte(key))
	if err != nil {
		c.JsonFail(content, http.StatusInternalServerError, 0, "密码加密失败", nil, "")
	}
	user = model.User{
		ID:       0,
		Nickname: req.Nickname,
		Password: string(encryptedPassword),
		Key:      key,
	}
	err = global.Db.Create(user).Error
	if err != nil {
		c.JsonFail(content, http.StatusInternalServerError, 0, "用户创建失败", nil, "")
	} else {
		c.JsonSuccess(content, nil, "")
	}
}

func (c *UserController) Login(content *gin.Context) {
	var req request.UserLoginRequest
	if err := content.ShouldBindJSON(&req); err != nil {
		c.JsonFail(content, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	c.JsonSuccess(content, nil, "")
}
