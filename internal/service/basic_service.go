package service

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"runtime"
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

func (s *BasicService) GetFunName()string {
	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
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