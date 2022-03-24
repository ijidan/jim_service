// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_model

import "time"

const TableNameDeviceAck = "device_ack"

// DeviceAck mapped from table <device_ack>
type DeviceAck struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 自增主键
	DeviceID   int64     `gorm:"column:device_id;not null" json:"device_id"`                               // 设备id
	Ack        int64     `gorm:"column:ack;not null;default:0" json:"ack"`                                 // 收到消息确认号
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	Device     Device    `gorm:"" json:"device"`
}

// TableName DeviceAck's table name
func (*DeviceAck) TableName() string {
	return TableNameDeviceAck
}
