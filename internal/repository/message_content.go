package repository

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"jim_service/internal/gen_model"
	"jim_service/internal/model"
	"time"
)

const (
	MessageIndexStatusToSend    = 1
	MessageIndexStatusSent      = 2
	MessageIndexStatusConfirmed = 3
	MessageIndexStatusDelete    = 0
)
type MessageContent struct {
	SenderId   uint64 `json:"sender_id,omitempty"`
	ReceiverId       uint64 `json:"receiver_id,omitempty"`
	GroupId          uint64 `json:"group_id,omitempty"`
	MessageContentId int64  `json:"message_content_id,omitempty"`
	MessageType      int32  `json:"message_type,omitempty"`
	MessageBody  []byte `json:"message_body,omitempty"`
	MessageExtra []byte `json:"message_extra,omitempty"`
}

func CreateMessageContent(db *gorm.DB, messageContent MessageContent) (*model.MessageContent, error) {
	index := &model.MessageIndex{
		SenderID:         messageContent.SenderId,
		ReceiverID:       messageContent.ReceiverId,
		MessageContentID: messageContent.MessageContentId,
		GroupID:          messageContent.GroupId,
		Status:           MessageIndexStatusToSend,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
		MessageContent:   gen_model.MessageContent{},
	}
	content := &model.MessageContent{
		ID:    messageContent.MessageContentId,
		Type:  messageContent.MessageType,
		Body:  string(messageContent.MessageBody),
		Extra: string(messageContent.MessageExtra),
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		err = db.Create(content).Error
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		err = db.Create(index).Error
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return content, nil
}
