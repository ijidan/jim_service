package gormt

import (
	"time"
)

// Group 群组
type Group struct {
	ID           uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                  // 自增主键
	Name         string    `gorm:"column:name;type:varchar(50);not null;default:''" json:"name"`                  // 群组名称
	Introduction string    `gorm:"column:introduction;type:varchar(255);not null;default:''" json:"introduction"` // 群组简介
	Extra        string    `gorm:"column:extra;type:varchar(1024);not null;default:''" json:"extra"`              // 附加属性
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`                              // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`                              // 更新时间
	DeletedAt    time.Time `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`                              // 删除时间
}

// GroupColumns get sql column name.获取数据库列名
var GroupColumns = struct {
	ID           string
	Name         string
	Introduction string
	Extra        string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "id",
	Name:         "name",
	Introduction: "introduction",
	Extra:        "extra",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}
