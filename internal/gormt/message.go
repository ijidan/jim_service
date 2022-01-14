package gormt

import (
	"time"
)

// Message 消息
type Message struct {
	ID           uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	SenderID     uint64    `gorm:"column:sender_id;type:bigint unsigned;not null;default:0" json:"senderId"`     // 发送人ID
	SenderType   int8      `gorm:"column:sender_type;type:tinyint;not null;default:0" json:"senderType"`         // 发送人类型
	ReceiverID   uint64    `gorm:"column:receiver_id;type:bigint unsigned;not null;default:0" json:"receiverId"` // 接收人ID，单聊则为user_id,群聊则为group_id
	ReceiverType int8      `gorm:"column:receiver_type;type:tinyint;not null;default:0" json:"receiverType"`     // 接收人类型
	AtUserID     string    `gorm:"column:at_user_id;type:varchar(255);not null;default:''" json:"atUserId"`      // 需要@的用户，多个用,分割
	MessageType  int       `gorm:"column:message_type;type:int;not null;default:0" json:"messageType"`           // 消息类型
	Status       int8      `gorm:"column:status;type:tinyint;not null;default:1" json:"status"`                  // 状态：1-正常，2-撤回，0-删除
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`                             // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`                             // 更新时间
	DeletedAt    time.Time `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`                             // 删除时间
}

// MessageColumns get sql column name.获取数据库列名
var MessageColumns = struct {
	ID           string
	SenderID     string
	SenderType   string
	ReceiverID   string
	ReceiverType string
	AtUserID     string
	MessageType  string
	Status       string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "id",
	SenderID:     "sender_id",
	SenderType:   "sender_type",
	ReceiverID:   "receiver_id",
	ReceiverType: "receiver_type",
	AtUserID:     "at_user_id",
	MessageType:  "message_type",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}
