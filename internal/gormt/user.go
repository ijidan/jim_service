package gormt

import (
	"time"
)

// User 用户
type User struct {
	ID        uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`              // 自增主键
	Nickname  string    `gorm:"column:nickname;type:varchar(20);not null;default:''" json:"nickname"`      // 昵称
	Password  string    `gorm:"column:password;type:varchar(100);not null;default:''" json:"password"`     // 密码
	Gender    string    `gorm:"column:gender;type:enum('1','2');not null" json:"gender"`                   // 性别，0:未知；1:男；2:女
	AvatarURL string    `gorm:"column:avatar_url;type:varchar(1024);not null;default:''" json:"avatarUrl"` // 用户头像链接
	Extra     string    `gorm:"column:extra;type:varchar(1024);not null;default:''" json:"extra"`          // 附加属性
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`                          // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`                          // 更新时间
	DeletedAt time.Time `gorm:"index:deleted_at;column:deleted_at;type:datetime" json:"deletedAt"`         // 删除时间
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID        string
	Nickname  string
	Password  string
	Gender    string
	AvatarURL string
	Extra     string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Nickname:  "nickname",
	Password:  "password",
	Gender:    "gender",
	AvatarURL: "avatar_url",
	Extra:     "extra",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}
