package gormt

import (
	"time"
)

// GroupUser 群组成员关系
type GroupUser struct {
	ID           uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                                                 // 自增主键
	GroupID      uint64    `gorm:"uniqueIndex:uk_group_id_user_id;column:group_id;type:bigint unsigned;not null" json:"groupId"`                 // 组id
	UserID       uint64    `gorm:"uniqueIndex:uk_group_id_user_id;index:idx_user_id;column:user_id;type:bigint unsigned;not null" json:"userId"` // 用户id
	UserShowName string    `gorm:"column:user_show_name;type:varchar(20);not null" json:"userShowName"`                                          // 用户在群组的昵称
	Extra        string    `gorm:"column:extra;type:varchar(1024);not null" json:"extra"`                                                        // 附加属性
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`                                                             // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`                                                             // 更新时间
	DeletedAt    time.Time `gorm:"index:deleted_at;column:deleted_at;type:datetime" json:"deletedAt"`                                            // 删除时间
}

// GroupUserColumns get sql column name.获取数据库列名
var GroupUserColumns = struct {
	ID           string
	GroupID      string
	UserID       string
	UserShowName string
	Extra        string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "id",
	GroupID:      "group_id",
	UserID:       "user_id",
	UserShowName: "user_show_name",
	Extra:        "extra",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}
