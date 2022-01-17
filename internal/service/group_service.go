package service

import (
	"context"
	"jim_service/config"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/repository"
	"jim_service/pkg"
)

type GroupService struct {
	BasicService
	proto_build.UnimplementedGroupServiceServer
}

func (s *GroupService) CreateGroup(c context.Context, req *proto_build.CreateGroupRequest) (*proto_build.CreateGroupResponse, error) {
	userId := s.GetLoginUserId()
	group, err := repository.CreateGroup(pkg.Db, userId, req.Name, req.Introduction, req.AvatarUrl)
	if err != nil {
		return nil, err
	}
	protoGroup, err1 := repository.GetProtoGroupByGroupId(pkg.Db, group.ID)
	if err != nil {
		return nil, err1
	}
	rsp := &proto_build.CreateGroupResponse{Group: protoGroup}
	defer s.AddSpan(c, pkg.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *GroupService) UpdateGroup(c context.Context, req *proto_build.UpdateGroupRequest) (*proto_build.UpdateGroupResponse, error) {
	err := repository.UpdateGroup(pkg.Db, req.Id, req.Name, req.Introduction, req.AvatarUrl)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UpdateGroupResponse{}
	return rsp, nil
}

func (s *GroupService) GetGroup(c context.Context, req *proto_build.GetGroupRequest) (*proto_build.GetGroupResponse, error) {
	group, err := repository.GetProtoGroup(pkg.Db, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.GetGroupResponse{Group: group}
	return rsp, nil
}

func (s *GroupService) QueryGroup(c context.Context, req *proto_build.QueryGroupRequest) (*proto_build.QueryGroupResponse, error) {
	protoGroupList, pager, err := repository.QueryProtoGroup(pkg.Db, req.GetKeyword(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.QueryGroupResponse{
		Pager: pager,
		Group: protoGroupList,
	}
	return rsp, nil
}

func (s *GroupService) DeleteGroup(c context.Context, req *proto_build.DeleteGroupRequest) (*proto_build.DeleteGroupResponse, error) {
	err := repository.DeleteGroup(pkg.Db, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.DeleteGroupResponse{}
	return rsp, nil
}

func (s *GroupService) JoinGroup(c context.Context, req *proto_build.JoinGroupRequest) (*proto_build.JoinGroupResponse, error) {
	err := repository.JoinGroup(pkg.Db, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.JoinGroupResponse{}
	return rsp, nil
}

func (s *GroupService) QuitGroup(c context.Context, req *proto_build.QuitGroupRequest) (*proto_build.QuitGroupResponse, error) {
	err := repository.QuitGroup(pkg.Db, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.QuitGroupResponse{}
	return rsp, nil
}

func NewGroupService(config *config.Config) *GroupService {
	instance := &GroupService{BasicService: BasicService{
		Name: "service_group",
		Host: config.Rpc.Host,
		Port: config.Rpc.Port,
		Ttl:  config.Rpc.Ttl,
	}}
	return instance
}
