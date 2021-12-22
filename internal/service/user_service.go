package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"jim_service/config"
	"jim_service/internal/jim_proto/proto_build"
)

// UserService Hello服务
type UserService struct {
	BasicService
	proto_build.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(c context.Context, req *proto_build.CreateUserRequest) (*proto_build.CreateUserResponse, error) {
	rsp := proto_build.CreateUserResponse{
		User: nil,
	}
	span, _ := opentracing.StartSpanFromContext(c, s.GetName())
	defer func() {
		span.SetTag("request", req)
		span.SetTag("reply", rsp.String())
		span.Finish()
	}()
	return &rsp, nil
}

// NewUserService 获取实例
func NewUserService(config *config.Config) *UserService {
	instance := &UserService{BasicService:BasicService{
		Name:    "user_service",
		AppName: config.App.Name,
		Host:    config.Rpc.Host,
		Port:    config.Rpc.Port,
		Ttl:     config.Rpc.Ttl,
	}}
	return instance
}
