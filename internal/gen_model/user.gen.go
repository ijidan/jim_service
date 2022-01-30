// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 自增主键
	Nickname  string         `gorm:"column:nickname;not null" json:"nickname"`          // 昵称
	Password  string         `gorm:"column:password;not null" json:"password"`          // 密码
	Gender    string         `gorm:"column:gender;not null" json:"gender"`              // 性别，0:未知；1:男；2:女
	AvatarURL string         `gorm:"column:avatar_url;not null" json:"avatar_url"`      // 用户头像链接
	Extra     string         `gorm:"column:extra;not null" json:"extra"`                // 附加属性
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`               // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`               // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`               // 删除时间
	Device    []Device       `gorm:"" json:"device"`
	Message   []Device       `gorm:"" json:"message"`
	GroupUser []GroupUser    `gorm:"" json:"group_user"`
	Feed      []Feed         `gorm:"" json:"feed"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
