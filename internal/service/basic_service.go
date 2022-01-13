package service

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"jim_service/internal/jim_proto/proto_build"
)


// BasicService 基本服务类
type BasicService struct {
	Name string
	Host string
	Port uint
	Ttl int64
}

// GetName 获取服务名称
func (s *BasicService) GetName() string {
	return s.Name
}

func (s *BasicService) GetAddress() string {
	address := fmt.Sprintf("%s:%d",s.Host, s.Port)
	return address
}

func (s *BasicService) GetTTL() int64 {
	return s.Ttl
}


func (s *BasicService)AddSpan(c context.Context,funcName string,req interface{},rsp interface{})  {
	span, _ := opentracing.StartSpanFromContext(c, s.Name)
	defer func() {
		span.SetTag("func",funcName)
		span.SetTag("request", req)
		span.SetTag("reply", rsp)
		span.Finish()
	}()
}

func (s *BasicService) SuccessCommon() *proto_build.CommonResponse  {
	rsp:=&proto_build.CommonResponse{
		ErrorCode:    0,
		BusinessCode: 0,
		Message:      "success",
	}
	return rsp
}

func (s *BasicService) FailCommon(errorCode uint64,businessCode uint64,message string) *proto_build.CommonResponse {
	rsp:=&proto_build.CommonResponse{
		ErrorCode:   errorCode ,
		BusinessCode: businessCode,
		Message:      message,
	}
	return rsp
}

func (s *BasicService) GetLoginUserId() uint64  {
	return 9
}