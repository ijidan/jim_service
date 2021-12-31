package gormt

import (
	"time"
)

// Device 设备
type Device struct {
	ID            uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                                  // 自增主键
	DeviceID      int64     `gorm:"unique;column:device_id;type:bigint;not null" json:"deviceId"`                                  // 设备id
	AppID         uint64    `gorm:"index:idx_app_id_user_id;column:app_id;type:bigint unsigned;not null" json:"appId"`             // app_id
	UserID        uint64    `gorm:"index:idx_app_id_user_id;column:user_id;type:bigint unsigned;not null;default:0" json:"userId"` // 账户id
	Type          int8      `gorm:"column:type;type:tinyint;not null" json:"type"`                                                 // 设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
	Brand         string    `gorm:"column:brand;type:varchar(20);not null" json:"brand"`                                           // 手机厂商
	Model         string    `gorm:"column:model;type:varchar(20);not null" json:"model"`                                           // 机型
	SystemVersion string    `gorm:"column:system_version;type:varchar(10);not null" json:"systemVersion"`                          // 系统版本
	SdkVersion    string    `gorm:"column:sdk_version;type:varchar(10);not null" json:"sdkVersion"`                                // app版本
	Status        int8      `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`                                   // 在线状态，0：离线；1：在线
	ConnAddr      string    `gorm:"column:conn_addr;type:varchar(25);not null" json:"connAddr"`                                    // 连接层服务器地址
	ConnFd        int64     `gorm:"column:conn_fd;type:bigint;not null" json:"connFd"`                                             // TCP连接对应的文件描述符
	CreateTime    time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`         // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"`         // 更新时间
}

// DeviceColumns get sql column name.获取数据库列名
var DeviceColumns = struct {
	ID            string
	DeviceID      string
	AppID         string
	UserID        string
	Type          string
	Brand         string
	Model         string
	SystemVersion string
	SdkVersion    string
	Status        string
	ConnAddr      string
	ConnFd        string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	DeviceID:      "device_id",
	AppID:         "app_id",
	UserID:        "user_id",
	Type:          "type",
	Brand:         "brand",
	Model:         "model",
	SystemVersion: "system_version",
	SdkVersion:    "sdk_version",
	Status:        "status",
	ConnAddr:      "conn_addr",
	ConnFd:        "conn_fd",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}
