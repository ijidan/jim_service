// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFeedLike = "feed_like"

// FeedLike mapped from table <feed_like>
type FeedLike struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement:false;default:0" json:"id"` // ID
	FeedID    uint64         `gorm:"column:feed_id;not null;default:0" json:"feed_id"`             // 动态ID
	UserID    uint64         `gorm:"column:user_id;not null;default:0" json:"user_id"`             // 用户ID
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName FeedLike's table name
func (*FeedLike) TableName() string {
	return TableNameFeedLike
}
