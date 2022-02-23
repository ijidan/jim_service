package gormt

import (
	"time"
)

// DeviceAck 设备消息同步序列号
type DeviceAck struct {
	ID         uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                          // 自增主键
	DeviceID   uint64    `gorm:"unique;column:device_id;type:bigint unsigned;not null" json:"deviceId"`                 // 设备id
	Ack        uint64    `gorm:"column:ack;type:bigint unsigned;not null;default:0" json:"ack"`                         // 收到消息确认号
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` // 更新时间
}

// DeviceAckColumns get sql column name.获取数据库列名
var DeviceAckColumns = struct {
	ID         string
	DeviceID   string
	Ack        string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	DeviceID:   "device_id",
	Ack:        "ack",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}
