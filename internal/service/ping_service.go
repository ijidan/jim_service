package service

import (
	"context"
	"jim_service/config"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
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
	defer s.AddSpan(c,pkg.GetFuncName(),req,rsp.String())
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
