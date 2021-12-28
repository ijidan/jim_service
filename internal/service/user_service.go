package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/types/known/timestamppb"
	"jim_service/config"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/model_gormt"
	"jim_service/pkg"
)

// UserService Hello服务
type UserService struct {
	BasicService
	proto_build.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(c context.Context, req *proto_build.CreateUserRequest) (*proto_build.CreateUserResponse, error) {
	gender:=proto_build.Gender_value[cast.ToString(req.GetGender())]
	userMgr:=model_gormt.UserMgr(pkg.Db)
	user:=&model_gormt.User{
		Nickname:  req.GetNickname(),
		Password:  req.GetPassword(),
		Key:       "",
		Gender:   cast.ToString(gender),
		AvatarURL: req.GetAvatarUrl(),
		Extra:     "",
	}
	userMgr.Create(user)

	rspUser:=&proto_build.User{
		Id:        user.ID,
		Nickname:  req.GetNickname(),
		Gender:    req.GetGender(),
		AvatarUrl: req.AvatarUrl,
		Extra:     user.Extra,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		DeletedAt: timestamppb.New(user.DeletedAt),
	}
	rsp := proto_build.CreateUserResponse{
		User: rspUser,
	}

	s.AddSpan(c,s.GetFunName(),req,rsp.String())
	return &rsp, nil
}

func (s *UserService)GetUser(c context.Context,req *proto_build.GetUserRequest)(*proto_build.GetUserResponse,error)  {
	id:=req.GetId()
	userMgr:=model_gormt.UserMgr(pkg.Db)
	user,err:=userMgr.GetFromID(id)
	if err!=nil{
		panic(err)
	}
	pbUser:=&proto_build.User{
		Id:        user.ID,
		Nickname:  user.Nickname,
		Gender:    proto_build.Gender_Male,
		AvatarUrl: user.AvatarURL,
		Extra:     user.Extra,
		CreatedAt: &timestamp.Timestamp{Seconds: user.CreatedAt.Unix()},
		UpdatedAt: &timestamp.Timestamp{Seconds: user.UpdatedAt.Unix()},
		DeletedAt: &timestamp.Timestamp{Seconds: user.DeletedAt.Unix()},
	}
	rsp:=&proto_build.GetUserResponse{
		User: pbUser,
	}
	return rsp,nil
}
// NewUserService 获取实例
func NewUserService(config *config.Config) *UserService {
	instance := &UserService{BasicService:BasicService{
		Name:    "service_ping",
		Host:    config.Rpc.Host,
		Port:    config.Rpc.Port,
		Ttl:     config.Rpc.Ttl,
	}}
	return instance
}
