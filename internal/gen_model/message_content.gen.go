// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMessageContent = "message_content"

// MessageContent mapped from table <message_content>
type MessageContent struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Type      int32          `gorm:"column:type;not null;default:0" json:"type"` // 消息类型
	Body      string         `gorm:"column:body;not null" json:"body"`           // 内容
	Extra     string         `gorm:"column:extra;not null" json:"extra"`         // 扩展
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`        // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`        // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`        // 删除时间
}

// TableName MessageContent's table name
func (*MessageContent) TableName() string {
	return TableNameMessageContent
}
