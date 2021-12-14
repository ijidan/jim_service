package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"jim/internal/rpc/proto_build"
)

// UserService Hello服务
type UserService struct {
	BasicService
	proto_build.UnimplementedUserServiceServer
}

// GetName 获取服务名称
func (s *UserService) GetName() string {
	return "user_service"
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
func NewUserService() *UserService {
	instance := &UserService{}
	return instance
}
