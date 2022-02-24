package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

			var data json.RawMessage
			msgData := dispatch.ClientMessage{
				Data: &data,
			}
			if err3 := json.Unmarshal(req.Data, &msgData); err3 != nil {
				color.Red("parse message err:%s", err3.Error())
				continue
			}
			var senderId string
			var receiverId string
			switch msgData.Type {
			case dispatch.MessageTypeText:
				var msgContent dispatch.TextMessage
				if err := json.Unmarshal(data, &msgContent); err != nil {
				}
				senderId = msgContent.SenderId
				receiverId = msgContent.ReceiverId
				//todo trie
			case dispatch.MessageTypeLocation:
				var msgContent dispatch.LocationMessage
				if err := json.Unmarshal(data, &msgContent); err != nil {
				}
				senderId = msgContent.SenderId
				receiverId = msgContent.ReceiverId
			case dispatch.MessageTypeFace:
				var msgContent dispatch.FaceMessage
				if err := json.Unmarshal(data, &msgContent); err != nil {
				}
				senderId = msgContent.SenderId
				receiverId = msgContent.ReceiverId
			case dispatch.MessageTypeSound:
				var msgContent dispatch.SoundMessage
				if err := json.Unmarshal(data, &msgContent); err != nil {
				}
				senderId = msgContent.SenderId
				receiverId = msgContent.ReceiverId
			case dispatch.MessageTypeImage:
				var msgContent dispatch.ImageMessage
				if err := json.Unmarshal(data, &msgContent); err != nil {
				}
				senderId = msgContent.SenderId
				receiverId = msgContent.ReceiverId
			case dispatch.MessageTypeFile:
				var msgContent dispatch.FileMessage
				if err := json.Unmarshal(data, &msgContent); err != nil {
				}
				senderId = msgContent.SenderId
				receiverId = msgContent.ReceiverId
			case dispatch.MessageTypeVideo:
				var msgContent dispatch.VideoMessage
				if err := json.Unmarshal(data, &msgContent); err != nil {
				}
				senderId = msgContent.SenderId
				receiverId = msgContent.ReceiverId
			}
			messageId := dispatch.GenMessageID()
			go func(senderId string, requestId uint32, messageId int64) {
				errAck := SendAckToSender(senderId, req.RequestId, messageId)
				if errAck != nil {
					color.Red("send ack error:%s", errAck.Error())
				}
			}(senderId, req.RequestId, messageId)

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

func NewGatewayService(config *config.Config) *GatewayService {
	instance := &GatewayService{BasicService: BasicService{
		Name: "service_gateway",
		Host: config.Rpc.Host,
		Port: config.Rpc.Port,
		Ttl:  config.Rpc.Ttl,
	}}
	return instance
}
