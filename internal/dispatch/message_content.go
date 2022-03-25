package dispatch

import (
	"encoding/json"
	"github.com/fatih/color"
	"jim_service/internal/jim_proto/proto_build"
)

func ParseMessage(req *proto_build.SendMessageRequest) (string, string, uint64, int64, string, []byte, []byte, error) {
	var data json.RawMessage
	msgData := ClientMessage{
		Data: &data,
	}
	if err3 := json.Unmarshal(req.Data, &msgData); err3 != nil {
		color.Red("parse message err:%s", err3.Error())
		return "", "", 0, 0, "", nil, nil, err3
	}
	var senderId string
	var receiverId string
	switch msgData.Type {
	case MessageTypeText:
		var msgContent TextMessage
		if err := json.Unmarshal(data, &msgContent); err != nil {
		}
		senderId = msgContent.SenderId
		receiverId = msgContent.ReceiverId
		//todo trie
	case MessageTypeLocation:
		var msgContent LocationMessage
		if err := json.Unmarshal(data, &msgContent); err != nil {
		}
		senderId = msgContent.SenderId
		receiverId = msgContent.ReceiverId
	case MessageTypeFace:
		var msgContent FaceMessage
		if err := json.Unmarshal(data, &msgContent); err != nil {
		}
		senderId = msgContent.SenderId
		receiverId = msgContent.ReceiverId
	case MessageTypeSound:
		var msgContent SoundMessage
		if err := json.Unmarshal(data, &msgContent); err != nil {
		}
		senderId = msgContent.SenderId
		receiverId = msgContent.ReceiverId
	case MessageTypeImage:
		var msgContent ImageMessage
		if err := json.Unmarshal(data, &msgContent); err != nil {
		}
		senderId = msgContent.SenderId
		receiverId = msgContent.ReceiverId
	case MessageTypeFile:
		var msgContent FileMessage
		if err := json.Unmarshal(data, &msgContent); err != nil {
		}
		senderId = msgContent.SenderId
		receiverId = msgContent.ReceiverId
	case MessageTypeVideo:
		var msgContent VideoMessage
		if err := json.Unmarshal(data, &msgContent); err != nil {
		}
		senderId = msgContent.SenderId
		receiverId = msgContent.ReceiverId
	}
	messageId := GenMessageID()

	return senderId, receiverId, 0, messageId, msgData.Type, req.Data, nil, nil

}
