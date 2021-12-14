package pkg

import "github.com/pkg/errors"

const (
	OK                 = 0    // Success
	NotLoggedIn        = 1000 // 未登录
	ParameterIllegal   = 1001 // 参数不合法
	UnauthorizedUserId = 1002 // 非法的用户Id
	Unauthorized       = 1003 // 未授权
	ServerError        = 1004 // 系统错误
	RoutingNotExist    = 1005 // 路由不存在
)

var errorMap = map[int]error{
	OK:                 errors.New("success"),
	NotLoggedIn:        errors.New("not login"),
	ParameterIllegal:   errors.New("param illegal"),
	UnauthorizedUserId: errors.New("UnauthorizedUserId"),
	Unauthorized:       errors.New("Unauthorized"),
	ServerError:        errors.New("server error"),
	RoutingNotExist:    errors.New("routing"),
}

// GetErrorMessage 根据错误码获取错误信息
func GetErrorMessage(code int) string {
	var message string
	if value, ok := errorMap[code]; ok {
		message = value.Error()
	} else {
		message = "error code not found"
	}
	return message
}
