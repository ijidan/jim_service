package dispatch

import (
	"github.com/bwmarrin/snowflake"
)

const (
	MessageTypeText     = "text"
	MessageTypeLocation = "location"
	MessageTypeFace     = "face"
	MessageTypeSound    = "sound"
	MessageTypeImage    = "image"
	MessageTypeFile     = "file"
	MessageTypeVideo    = "video"
	MessageTypeAck      = "ack"
)

var MessageTypeToDbType=map[string]int32{
	MessageTypeText     :1,
	MessageTypeLocation :2,
	MessageTypeFace     :3,
	MessageTypeSound    :4,
	MessageTypeImage    :5,
	MessageTypeFile     :6,
	MessageTypeVideo    :7,
	MessageTypeAck      :8,
}


type ClientMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type TextMessage struct {
	SenderId     string `json:"sender_id"`
	SenderName   string `json:"sender_name"`
	SenderAvatar string `json:"sender_avatar"`
	ReceiverId   string `json:"receiver_id"`
	Content      string `json:"content"`
}
type LocationMessage struct {
	SenderId     string  `json:"sender_id"`
	SenderName   string  `json:"sender_name"`
	SenderAvatar string  `json:"sender_avatar"`
	ReceiverId   string  `json:"receiver_id"`
	CoverImage   string  `json:"cover_image"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	MapLink      string  `json:"map_link"`
	Desc         string  `json:"desc"`
}

type FaceMessage struct {
	SenderId     string `json:"sender_id"`
	SenderName   string `json:"sender_name"`
	SenderAvatar string `json:"sender_avatar"`
	ReceiverId   string `json:"receiver_id"`
	Symbol       string `json:"symbol"`
}

type SoundMessage struct {
		SenderId     string `json:"sender_id"`
		SenderName   string `json:"sender_name"`
		SenderAvatar string `json:"sender_avatar"`
		ReceiverId   string `json:"receiver_id"`
		Url          string `json:"url"`
		Size         uint32    `json:"size"`
		Seconds      uint32    `json:"seconds"`
}

type ImageMessage struct {
	SenderId     string `json:"sender_id"`
	SenderName   string `json:"sender_name"`
	SenderAvatar string `json:"sender_avatar"`
	ReceiverId   string `json:"receiver_id"`
	Format       uint32    `json:"format"`
	Size         uint32    `json:"size "`
	Width        uint32    `json:"width "`
	Height       uint32    `json:"height "`
	IconUrl      string `json:"icon_url "`
	BigUrl       string `json:"big_url "`
}

type FileMessage struct {
	SenderId     string `json:"sender_id"`
	SenderName   string `json:"sender_name"`
	SenderAvatar string `json:"sender_avatar"`
	ReceiverId   string `json:"receiver_id"`
	Size         uint32    `json:"size"`
	Name         string `json:"name"`
	Format       uint32    `json:"format"`
	ThumbUrl     string `json:"thumb_url"`
	Url          string `json:"url"`
}
type VideoMessage struct {
	SenderId     string `json:"sender_id"`
	SenderName   string `json:"sender_name"`
	SenderAvatar string `json:"sender_avatar"`
	ReceiverId   string `json:"receiver_id"`
	Size         uint32    `json:"size"`
	Seconds      uint32    `json:"seconds"`
	Url          string `json:"url"`
	Format       uint32    `json:"format"`
	ThumbUrl     string `json:"thumb_url"`
	ThumbSize    uint32    `json:"thumb_size"`
	ThumbWidth   uint32    `json:"thumb_width"`
	ThumbHeight  uint32    `json:"thumb_height"`
	ThumbFormat  uint32    `json:"thumb_format"`
}


type AckMessage struct {
		ReceiverId string `json:"receiver_id"`
		RequestId  uint32    `json:"request_id"`
		MessageId  int64    `json:"message_id"`
}

func GenMessageID() int64 {
	node, _ := snowflake.NewNode(1)
	return node.Generate().Int64()
}
