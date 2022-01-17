package repository

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/model"
)

const (
	SenderTypeUser  = 1
	SenderTypeGroup = 2

	ReceiverTypeUser  = 1
	ReceiverTypeGroup = 2

	MessageTypeText     = 0
	MessageTypeLocation = 1
	MessageTypeFace     = 2
	MessageTypeSound    = 3
	MessageTypeImage    = 4
	MessageTypeFile     = 5
	MessageTypeVideo    = 6
)

func CreateUserTextMessage(db *gorm.DB, senderId uint64, toUserId uint64, text *proto_build.TextMessage) (*model.Message, error) {
	message := &model.Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeText,
		Data:         text.Content,
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserLocationMessage(db *gorm.DB, senderId uint64, toUserId uint64, location *proto_build.LocationMessage) (*model.Message, error) {
	message := &model.Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeLocation,
		Data:         location.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserFaceMessage(db *gorm.DB, senderId uint64, toUserId uint64, face *proto_build.FaceMessage) (*model.Message, error) {
	message := &model.Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeFace,
		Data:         face.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserSoundMessage(db *gorm.DB, senderId uint64, toUserId uint64, sound *proto_build.SoundMessage) (*model.Message, error) {
	message := &model.Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeSound,
		Data:         sound.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserVideoMessage(db *gorm.DB, senderId uint64, toUserId uint64, video *proto_build.VideoMessage) (*model.Message, error) {
	message := &model.Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeVideo,
		Data:         video.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserImageMessage(db *gorm.DB, senderId uint64, toUserId uint64, image *proto_build.ImageMessageItem) (*model.Message, error) {
	message := &model.Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeImage,
		Data:         image.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserFileMessage(db *gorm.DB, senderId uint64, toUserId uint64, file *proto_build.FileMessage) (*model.Message, error) {
	message := &model.Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeFile,
		Data:         file.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}