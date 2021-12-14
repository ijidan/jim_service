package pkg

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type Response struct {
}

func (r *Response) JsonSuccess(data map[string]interface{}, jumpUrl string) gin.H {
	return gin.H{"code": 0, "message": "success", "data": data, "jump_url": jumpUrl}
}

func (r *Response) JsonFail(code int32, businessCode int32, message string, data map[string]interface{}, jumpUrl string) gin.H {
	return gin.H{"code": code, "business_code": businessCode, "message": message, "data": data, "jump_url": jumpUrl}
}

func (r *Response) Pager() {
}

var (
	onceResponse     sync.Once
	instanceResponse *Response
)

func GetResponseInstance() *Response {
	onceResponse.Do(func() {
		instanceResponse = &Response{}
	})
	return instanceResponse
}
