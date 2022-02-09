package service

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"jim_service/config"
	"jim_service/internal/dispatch"
	"jim_service/internal/jim_proto/proto_build"
)

type GatewayService struct {
	BasicService
	proto_build.UnimplementedGatewayServiceServer
}

func (s *GatewayService) Register(c context.Context, req *proto_build.RegisterRequest) (*proto_build.RegisterResponse, error) {

	dispatch.ClientIdGatewayIdMap.Store(req.ClientId,req.GatewayId)
	color.Cyan("grpc client register success,client id:%s,gateway id:%s",req.ClientId,req.GatewayId)
	rsp:=&proto_build.RegisterResponse{}
	return rsp, nil
}

func (s *GatewayService) UnRegister(c context.Context, req *proto_build.UnRegisterRequest) (*proto_build.UnRegisterResponse, error) {
	dispatch.ClientIdGatewayIdMap.Delete(req.GatewayId)
	color.Cyan("grpc client unregister success,client id:%s,gateway id:%s",req.ClientId,req.GatewayId)
	rsp:=&proto_build.UnRegisterResponse{}
	return rsp,nil
}

func (s *GatewayService) SendMessage(serv proto_build.GatewayService_SendMessageServer) error {
	errCh:= make(chan error)
	go func(errCh chan error) {
		for  {
			req,err:=serv.Recv()
			if err!=nil{
				errCh <- err
				return
			}
			dispatch.GatewayIdSendMessageMap.Store(req.GatewayId,serv)
			color.Cyan("%s",dispatch.Dump())
			color.Yellow("grpc received message:%s", string(req.Data))

			var data json.RawMessage
			msg:=dispatch.ClientMessage{
				Data: &data,
			}
			if string(req.Data)=="ping"{
				continue
			}
			if err3 := json.Unmarshal(req.Data, &msg); err3 != nil {
				color.Red("parse message err:%s",err3.Error())
			}
			switch msg.Cmd {
			case "chat.c2c.txt":
				content:=dispatch.TextMessage{}
				if err4 := json.Unmarshal(data, &content); err4 != nil {
					color.Red("parse message err:%s",err4.Error())
				}
				toReceiverId:=content.ToReceiverId
				gatewayId,ok1:=dispatch.ClientIdGatewayIdMap.Load(toReceiverId)
				if !ok1{
					color.Red("compute Gateway Id err")
				}
				serv,ok2:=dispatch.GatewayIdSendMessageMap.Load(gatewayId)
				if !ok2{
					color.Red("compute send message serv err")
				}
				rsp:=&proto_build.SendMessageResponse{Data: req.Data}
				srv:=serv.(proto_build.GatewayService_SendMessageServer)
				err5:=srv.Send(rsp)
				if err5!=nil{
					color.Red("trans message err:%s",err5.Error())
				}
			}


		}
	}(errCh)
	err:=<-errCh
	return status.Errorf(codes.Internal, err.Error())
}

func (s *GatewayService) PushToAllMessage(serv proto_build.GatewayService_PushToAllMessageServer) error {
	errCh:= make(chan error)
	go func(errCh chan error) {
		for  {
			req,err:=serv.Recv()
			if err!=nil{
				errCh <- err
				return
			}
			dispatch.GatewayIdPushToAllMap.Store(req.GatewayId,serv)
			color.Cyan("%s",dispatch.Dump())
			color.Yellow("grpc received push message:%s", string(req.Data))
		}
	}(errCh)
	err:=<-errCh
	return status.Errorf(codes.Internal, err.Error())
}

func NewGatewayService(config *config.Config) *GatewayService {
	instance := &GatewayService{BasicService:BasicService{
		Name:    "service_gateway",
		Host:    config.Rpc.Host,
		Port:    config.Rpc.Port,
		Ttl:     config.Rpc.Ttl,
	}}
	return instance
}
