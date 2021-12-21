package service

import (
	"fmt"
	"jim_service/global"
)

// BasicService 基本服务类
type BasicService struct {
	Name string
}

// GetName 获取服务名称
func (s *BasicService) GetName() string {
	return s.Name
}


func (s *BasicService) GetRegisterKey()string {
	target := fmt.Sprintf("/etcdv3://%s/grpc/%s", global.Config.App.Name, s.GetName())
	return target
}


