package service

import (
	"fmt"
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
