package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"jim_service/config"
	"jim_service/internal/jim_proto/proto_build"
)

// PingService Hello服务
type PingService struct {
	BasicService
	proto_build.UnimplementedPingServiceServer
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
func NewPingService(config *config.Config) *PingService {
	instance := &PingService{BasicService:BasicService{
		Name:    "service_ping",
		Host:    config.Rpc.Host,
		Port:    config.Rpc.Port,
		Ttl:     config.Rpc.Ttl,
	}}
	return instance
}
