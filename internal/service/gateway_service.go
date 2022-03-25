package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"jim_service/config"
	"jim_service/internal/dispatch"
	"jim_service/internal/jim_proto/proto_build"
	"time"
)

type GatewayService struct {
	BasicService
	proto_build.UnimplementedGatewayServiceServer
}

func (s *GatewayService) Register(c context.Context, req *proto_build.RegisterRequest) (*proto_build.RegisterResponse, error) {
	dispatch.ClientIdGatewayIdMap.Store(req.ClientId, req.GatewayId)
	color.Cyan("grpc client register success,client id:%s,gateway id:%s", req.ClientId, req.GatewayId)
	rsp := &proto_build.RegisterResponse{}
	return rsp, nil
}

func (s *GatewayService) UnRegister(c context.Context, req *proto_build.UnRegisterRequest) (*proto_build.UnRegisterResponse, error) {
	dispatch.ClientIdGatewayIdMap.Delete(req.GatewayId)
	color.Cyan("grpc client unregister success,client id:%s,gateway id:%s", req.ClientId, req.GatewayId)
	rsp := &proto_build.UnRegisterResponse{}
	return rsp, nil
}

func (s *GatewayService) SendToAll(c context.Context, req *proto_build.SendToAllRequest) (*proto_build.SendToAllResponse, error) {
	requestId := uint32(time.Now().Second())
	go SendToAll(dispatch.BusinessCmdS2C,requestId,req.Data)
	rsp:=&proto_build.SendToAllResponse{}
	return rsp,nil
}



func (s *GatewayService) SendMessage(stream proto_build.GatewayService_SendMessageServer) error {
	errCh := make(chan error)
	go func(errCh chan error) {
		for {
			req, err := stream.Recv()

			if err != nil {
				errCh <- err
				return
			}
			dispatch.GatewayIdSendMessageMap.Store(req.GatewayId, stream)
			color.Cyan("%s", dispatch.Dump())
			color.Yellow("grpc received message:%s", string(req.Data))

			if req.Cmd==dispatch.BusinessCmdPing{
				continue
			}
			senderId, receiverId, groupId, messageContentId, messageType, messageBody, messageExtra, err2 := dispatch.ParseMessage(req)
			if err2 != nil {
				color.Red("parse message err:%s", err2.Error())
				continue
			}

			go func(senderId string, requestId uint32, messageId int64) {
				errAck := SendAckToSender(senderId, req.RequestId, messageId)
				if errAck != nil {
					color.Red("send ack error:%s", errAck.Error())
				}
			}(senderId, req.RequestId, messageContentId)


			//send message to kafka
			go func(senderId uint64, receiverId uint64, groupId uint64, messageContentId int64, messageType string, messageBody []byte, messageExtra []byte) {
				dispatch.PublishMessage(senderId, receiverId, groupId, messageContentId, messageType, messageBody, messageExtra)
			}(cast.ToUint64(senderId), cast.ToUint64(receiverId), groupId, messageContentId, messageType, messageBody, messageExtra)

			//send ack
			go func(receiverId string, cmd string, requestId uint32, data []byte) {
				errSend := TransferToReceiver(receiverId, cmd, requestId, data)
				if errSend != nil {
					color.Red("send ack error:%s", errSend.Error())
				}
			}(receiverId, req.Cmd, req.RequestId, req.Data)
		}
	}(errCh)
	err := <-errCh
	return status.Errorf(codes.Internal, err.Error())
}

func SendAckToSender(senderId string, requestId uint32, messageId int64) error {
	gatewayId, ok1 := dispatch.ClientIdGatewayIdMap.Load(senderId)
	if !ok1 {
		msg := "compute Gateway Id err"
		return errors.New(msg)
	}
	serv, ok2 := dispatch.GatewayIdSendMessageMap.Load(gatewayId)
	if !ok2 {
		msg := "compute send message serv err"
		return errors.New(msg)
	}

	//replay ack
	ackMessage := dispatch.AckMessage{
		ReceiverId: senderId,
		RequestId:  requestId,
		MessageId:  messageId,
	}
	ackJson, _ := json.Marshal(ackMessage)
	rsp := &proto_build.SendMessageResponse{
		GatewayId:  gatewayId.(string),
		Cmd:        dispatch.BusinessCmdC2C,
		RequestId:  requestId,
		Data:       ackJson,
		ReceiverId: senderId,
	}
	srv := serv.(proto_build.GatewayService_SendMessageServer)
	err5 := srv.Send(rsp)
	if err5 != nil {
		msg := fmt.Sprintf("send ack message err:%s", err5.Error())
		return errors.New(msg)
	}
	return nil
}

func TransferToReceiver(receiverId string, cmd string, requestId uint32, data []byte) error {
	gatewayId, ok1 := dispatch.ClientIdGatewayIdMap.Load(receiverId)
	if !ok1 {
		msg := "compute Gateway Id err"
		return errors.New(msg)
	}
	serv, ok2 := dispatch.GatewayIdSendMessageMap.Load(gatewayId)
	if !ok2 {
		msg := "compute send message serv err"
		return errors.New(msg)
	}

	rsp := &proto_build.SendMessageResponse{
		GatewayId:  gatewayId.(string),
		Cmd:        cmd,
		RequestId:  requestId,
		Data:       data,
		ReceiverId: receiverId,
	}
	srv := serv.(proto_build.GatewayService_SendMessageServer)
	err5 := srv.Send(rsp)
	if err5 != nil {
		msg := fmt.Sprintf("send ack message err:%s", err5.Error())
		return errors.New(msg)
	}
	return nil
}

func SendToAll(cmd string, requestId uint32, data []byte) {
	color.Yellow("send to all begin")
	dispatch.ClientIdGatewayIdMap.Range(func(key, value interface{}) bool {
		receiverId := key.(string)
		gatewayId := value.(string)
		color.Yellow("send to all range: %s:%s",gatewayId,receiverId)
		rsp := &proto_build.SendMessageResponse{
			GatewayId:  gatewayId,
			Cmd:        cmd,
			RequestId:  requestId,
			Data:       data,
			ReceiverId: receiverId,
		}
		serv, ok := dispatch.GatewayIdSendMessageMap.Load(gatewayId)
		if ok {
			srv := serv.(proto_build.GatewayService_SendMessageServer)
			err := srv.Send(rsp)
			if err != nil {
				color.Red("send to all error,send message err:%s", err.Error())
			}else{
				color.Yellow("send to all success")
			}
		} else {
			color.Red("send to all error,gateway id:%s", gatewayId)
		}
		return true
	})
}

func NewGatewayService(config *config.Config) *GatewayService {
	instance := &GatewayService{BasicService: BasicService{
		Name: "service_gateway",
		Host: config.Rpc.Host,
		Port: config.Rpc.Port,
		Ttl:  config.Rpc.Ttl,
	}}
	return instance
}
