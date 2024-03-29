// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_model

import "time"

const TableNameGid = "gid"

// Gid mapped from table <gid>
type Gid struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 自增主键
	BusinessID  string    `gorm:"column:business_id;not null" json:"business_id"`                           // 业务id
	MaxID       int64     `gorm:"column:max_id;not null;default:0" json:"max_id"`                           // 最大id
	Step        int32     `gorm:"column:step;not null;default:1000" json:"step"`                            // 步长
	Description string    `gorm:"column:description;not null" json:"description"`                           // 描述
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
}

// TableName Gid's table name
func (*Gid) TableName() string {
	return TableNameGid
}
