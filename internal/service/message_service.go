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

type MessageService struct {
	BasicService
	proto_build.UnimplementedMessageServiceServer
}

func (s *MessageService) SendUserTextMessage(c context.Context, req *proto_build.SendUserTextMessageRequest) (*proto_build.SendUserTextMessageResponse, error) {
	userId:=s.GetLoginUserId()
	message,err:=repository.CreateUserTextMessage(pkg.Db,userId,req.ToUserId,req.Text)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	rsp:=&proto_build.SendUserTextMessageResponse{Id: message.ID}
	return rsp,nil
}

func (s *MessageService) SendUserLocationMessage(c context.Context, req *proto_build.SendUserLocationMessageRequest) (*proto_build.SendUserLocationMessageResponse, error) {
	userId:=s.GetLoginUserId()
	message,err:=repository.CreateUserLocationMessage(pkg.Db,userId,req.ToUserId,req.Location)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	rsp:=&proto_build.SendUserLocationMessageResponse{Id: message.ID}
	return rsp,nil
}

func (s *MessageService) SendUserFceMessage(c context.Context, req *proto_build.SendUserFaceMessageRequest) (*proto_build.SendUserFaceMessageResponse, error) {
	userId:=s.GetLoginUserId()
	message,err:=repository.CreateUserFaceMessage(pkg.Db,userId,req.ToUserId,req.Face)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	rsp:=&proto_build.SendUserFaceMessageResponse{Id: message.ID}
	return rsp,nil
}

func (s *MessageService) SendUserSoundMessage(c context.Context, req *proto_build.SendUserSoundMessageRequest) (*proto_build.SendUserSoundMessageResponse, error) {
	userId:=s.GetLoginUserId()
	message,err:=repository.CreateUserSoundMessage(pkg.Db,userId,req.ToUserId,req.Sound)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	rsp:=&proto_build.SendUserSoundMessageResponse{Id: message.ID}
	return rsp,nil
}

func (s *MessageService) SendUserVideoMessage(c context.Context, req *proto_build.SendUserVideoMessageRequest) (*proto_build.SendUserVideoMessageResponse, error) {
	userId:=s.GetLoginUserId()
	message,err:=repository.CreateUserVideoMessage(pkg.Db,userId,req.ToUserId,req.Video)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	rsp:=&proto_build.SendUserVideoMessageResponse{Id: message.ID}
	return rsp,nil
}

func (s *MessageService) SendUserImageMessage(c context.Context, req *proto_build.SendUserImageMessageRequest) (*proto_build.SendUserImageMessageResponse, error) {
	userId:=s.GetLoginUserId()
	message,err:=repository.CreateUserImageMessage(pkg.Db,userId,req.ToUserId,req.Image)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	rsp:=&proto_build.SendUserImageMessageResponse{Id: message.ID}
	return rsp,nil
}

func (s *MessageService) SendUserFileMessage(c context.Context, req *proto_build.SendUserFileMessageRequest) (*proto_build.SendUserFileMessageResponse, error) {
	userId:=s.GetLoginUserId()
	message,err:=repository.CreateUserFileMessage(pkg.Db,userId,req.ToUserId,req.File)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	rsp:=&proto_build.SendUserFileMessageResponse{Id: message.ID}
	return rsp,nil
}



func (s *MessageService) SendGroupTextMessage(c context.Context, req *proto_build.SendGroupTextMessageRequest) (*proto_build.SendGroupTextMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupTextMessage not implemented")
}

func (s *MessageService) SendGroupLocationMessage(c context.Context, req *proto_build.SendGroupLocationMessageRequest) (*proto_build.SendGroupLocationMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupLocationMessage not implemented")
}

func (s *MessageService) SendGroupFceMessage(c context.Context, req *proto_build.SendGroupFaceMessageRequest) (*proto_build.SendGroupFaceMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupFceMessage not implemented")
}

func (s *MessageService) SendGroupSoundMessage(c context.Context, req *proto_build.SendGroupSoundMessageRequest) (*proto_build.SendGroupSoundMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupSoundMessage not implemented")
}

func (s *MessageService) SendGroupVideoMessage(c context.Context, req *proto_build.SendGroupVideoMessageRequest) (*proto_build.SendGroupVideoMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupVideoMessage not implemented")
}

func (s *MessageService) SendGroupImageMessage(c context.Context, req *proto_build.SendGroupImageMessageRequest) (*proto_build.SendGroupImageMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupImageMessage not implemented")
}

func (s *MessageService) SendGroupFileMessage(c context.Context, req *proto_build.SendGroupFileMessageRequest) (*proto_build.SendGroupFileMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupFileMessage not implemented")
}

func NewMessageService(config *config.Config) *MessageService {
	instance := &MessageService{BasicService: BasicService{
		Name: "service_message",
		Host: config.Rpc.Host,
		Port: config.Rpc.Port,
		Ttl:  config.Rpc.Ttl,
	}}
	return instance
}
