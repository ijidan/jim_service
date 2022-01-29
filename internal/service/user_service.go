package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"jim_service/config"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/repository"
	"jim_service/pkg"
)

type UserService struct {
	BasicService
	proto_build.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(c context.Context, req *proto_build.UserCreateRequest) (*proto_build.UserCreateResponse, error) {
	if req.Password != req.PasswordRpt {
		return nil, status.Error(codes.Internal, "password ,password_rpt not the same")
	}
	gender := repository.GenderProtoToDb[req.Gender]
	user, err := repository.CreateUser(pkg.Db, req.Nickname, req.Password, gender, req.AvatarUrl, "")
	if err != nil {
		return nil, err
	}
	go repository.PublishNewUser(user.ID)
	protoUser, err1 := repository.GetProtoUserByUserId(pkg.Db, user.ID)
	if err != nil {
		return nil, err1
	}
	rsp := &proto_build.UserCreateResponse{User: protoUser}
	defer s.AddSpan(c, pkg.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *UserService) GetUser(c context.Context, req *proto_build.UserGetRequest) (*proto_build.UserGetResponse, error) {
	id := req.GetId()
	protoUser, err := repository.GetProtoUserByUserId(pkg.Db, id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UserGetResponse{
		User: protoUser,
	}
	defer s.AddSpan(c, pkg.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *UserService) QueryUser(c context.Context, req *proto_build.UserQueryRequest) (*proto_build.UserQueryResponse, error) {
	protoUserList, pager, err := repository.QueryProtoUser(pkg.Db, req.GetKeyword(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UserQueryResponse{
		Pager: pager,
		User:  protoUserList,
	}
	return rsp, nil
}

func (s *UserService) UserLogin(c context.Context, req *proto_build.UserLoginRequest) (*proto_build.UserLoginResponse, error) {
	protoUser, err := repository.GetProtoUser(pkg.Db, req.GetNickname(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	userId := protoUser.Id
	token := pkg.GenJwtToken(userId, pkg.Conf.Jwt.Secret)
	rsp := &proto_build.UserLoginResponse{Token: token}
	return rsp, nil
}

func (s *UserService) UpdatePassword(c context.Context, req *proto_build.UpdatePasswordRequest) (*proto_build.UpdatePasswordResponse, error) {
	if req.Password != req.PasswordRpt {
		return nil, status.Error(codes.Internal, "password ,password_rpt not the same")
	}
	userId := s.GetLoginUserId()
	err := repository.UpdateUserPassword(pkg.Db, userId, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.UpdatePasswordResponse{}
	return rsp, nil
}

func (s *UserService) UpdateAvatar(c context.Context, req *proto_build.UpdateAvatarRequest) (*proto_build.UpdateAvatarResponse, error) {
	url := req.GetUrl()
	userId := s.GetLoginUserId()
	err := repository.UpdateUserAvatar(pkg.Db, userId, url)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.UpdateAvatarResponse{}
	return rsp, nil
}

func NewUserService(config *config.Config) *UserService {
	instance := &UserService{BasicService: BasicService{
		Name: "service_ping",
		Host: config.Rpc.Host,
		Port: config.Rpc.Port,
		Ttl:  config.Rpc.Ttl,
	}}
	return instance
}
