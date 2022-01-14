// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGroup = "group"

// Group mapped from table <group>
type Group struct {
	ID           uint64         `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 自增主键
	Name         string         `gorm:"column:name;not null" json:"name"`                  // 群组名称
	Introduction string         `gorm:"column:introduction;not null" json:"introduction"`  // 群组简介
	Extra        string         `gorm:"column:extra;not null" json:"extra"`                // 附加属性
	UserID       int64          `gorm:"column:user_id;not null;default:0" json:"user_id"`  // 群主ID
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`               // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`               // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`               // 删除时间
	GroupUser    []GroupUser    `gorm:"" json:"group_user"`
}

// TableName Group's table name
func (*Group) TableName() string {
	return TableNameGroup
}
