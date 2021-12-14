package middleware

import (
	sContext "context"
	"github.com/gin-gonic/gin"
	"jim/global"
	"jim/pkg"
	"time"
)

func AppSetting() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(
			"app_name",
			global.Config.App.Name,
		)
		context.Set("app_ver", global.Config.App.Ver)
		requestId := context.Request.Header.Get(global.RequestId)
		if requestId == "" {
			requestId = pkg.GetUUId()
		}
		context.Set(global.RequestId, requestId)
		context.Writer.Header().Set(global.RequestId, requestId)

		ctx, cancel := sContext.WithTimeout(context, 2*time.Second)
		defer cancel()
		context.Set("cancel_ctx", ctx)

		context.Next()
	}
}
