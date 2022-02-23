package gormt

import (
	"time"
)

// Gid 分布式自增主键
type Gid struct {
	ID          uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                          // 自增主键
	BusinessID  string    `gorm:"unique;column:business_id;type:varchar(128);not null" json:"businessId"`                // 业务id
	MaxID       uint64    `gorm:"column:max_id;type:bigint unsigned;not null;default:0" json:"maxId"`                    // 最大id
	Step        uint      `gorm:"column:step;type:int unsigned;not null;default:1000" json:"step"`                       // 步长
	Description string    `gorm:"column:description;type:varchar(255);not null" json:"description"`                      // 描述
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` // 更新时间
}

// GidColumns get sql column name.获取数据库列名
var GidColumns = struct {
	ID          string
	BusinessID  string
	MaxID       string
	Step        string
	Description string
	CreateTime  string
	UpdateTime  string
}{
	ID:          "id",
	BusinessID:  "business_id",
	MaxID:       "max_id",
	Step:        "step",
	Description: "description",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}
