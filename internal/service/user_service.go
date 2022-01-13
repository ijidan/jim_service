package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"jim_service/config"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/repository"
	"jim_service/pkg"
)

type UserService struct {
	BasicService
	proto_build.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(c context.Context, req *proto_build.CreateUserRequest) (*proto_build.CreateUserResponse, error) {
	if req.Password != req.PasswordRpt {
		return nil, status.Error(codes.Internal, "password ,password_rpt not the same")
	}
	gender := repository.GenderProtoToDb[req.Gender]
	user, err := repository.CreateUser(pkg.Db, req.Nickname, req.Password, gender, req.AvatarUrl, "")
	if err != nil {
		return nil, err
	}
	protoUser, err1 := repository.GetProtoUserByUserId(pkg.Db, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err1.Error())
	}
	rsp := proto_build.CreateUserResponse{User: protoUser}
	defer s.AddSpan(c, pkg.GetFuncName(), req, rsp.String())
	return &rsp, nil
}

func (s *UserService) GetUser(c context.Context, req *proto_build.GetUserRequest) (*proto_build.GetUserResponse, error) {
	id := req.GetId()
	protoUser, err := repository.GetProtoUserByUserId(pkg.Db, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.GetUserResponse{
		User: protoUser,
	}
	defer s.AddSpan(c, pkg.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *UserService) QueryUser(c context.Context, req *proto_build.QueryUserRequest) (*proto_build.QueryUserResponse, error) {
	protoUserList, pager, err := repository.QueryProtoUser(pkg.Db, req.GetKeyword(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.QueryUserResponse{
		Pager: pager,
		User:  protoUserList,
	}
	return rsp, nil
}

func (s *UserService) UpdateAvatar(stream proto_build.UserService_UpdateAvatarServer) error {
	fileUpload:=pkg.NewFileUpload()
	defer fileUpload.Close()

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return status.Error(codes.Internal, err.Error())
		}
		content := req.GetContent()
		_,err1:=fileUpload.WriteContent(content)
		if err1 != nil {
			return status.Error(codes.Internal, "write file failed")
		}
	}

	fileName:=fileUpload.GetFileName()
	img,err2:=pkg.SmMsUploadImage(pkg.Conf.Smms.User,pkg.Conf.Smms.Password,fileName)
	if err2!=nil{
		return status.Error(codes.Internal,err2.Error())
	}
	rsp := &proto_build.UploadAvatarResponse{Url: img.Url}
	err3 := stream.SendAndClose(rsp)
	if err3 != nil {
		return status.Error(codes.Internal, err3.Error())
	}
	return nil
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
