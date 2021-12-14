package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"jim/global"
	"jim/pkg"
	"net/http"
	"time"
)

func JwtVerify() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		if token == "" {
			rsp := global.Response.JsonFail(pkg.Unauthorized, pkg.OK, "token required", nil, "")
			context.JSON(http.StatusOK, rsp)
		}
		claim, err := pkg.ParseJwtToken(token, global.Config.Jwt.Secret)
		if err != nil {
			rsp := global.Response.JsonFail(pkg.Unauthorized, pkg.OK, "token error", nil, "")
			context.JSON(http.StatusOK, rsp)
		}

		exp := cast.ToInt(claim["exp"])
		if exp < time.Now().Second() {
			rsp := global.Response.JsonFail(pkg.Unauthorized, pkg.OK, "token expired", nil, "")
			context.JSON(http.StatusOK, rsp)
		}
		uid := cast.ToInt(claim["uid"])
		context.Set("user_id", uid)

		context.Next()
	}
}
