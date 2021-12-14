package service

// BasicService 基本服务类
type BasicService struct {
	Name string
}

// GetName 获取服务名称
func (s *BasicService) GetName() string {
	return s.Name
}
