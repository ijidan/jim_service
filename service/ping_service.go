package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"jim/internal/rpc/proto_build"
)

// PingService Hello服务
type PingService struct {
	BasicService
	proto_build.UnimplementedPingServiceServer
}

// GetName 获取服务名称
func (s *PingService) GetName() string {
	return "ping_service"
}

func (s *PingService) Ping(c context.Context, req *proto_build.PingRequest) (*proto_build.PingResponse, error) {
	rsp := &proto_build.PingResponse{
		Message: "pong",
	}
	span, _ := opentracing.StartSpanFromContext(c, s.GetName())
	defer func() {
		span.SetTag("request", req)
		span.SetTag("reply", rsp.String())
		span.Finish()
	}()
	return rsp, nil
}

// NewPingService 获取实例
func NewPingService() *PingService {
	instance := &PingService{}
	return instance
}
