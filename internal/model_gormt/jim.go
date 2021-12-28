package model_gormt

import (
	"time"
)

// Device 设备
type Device struct {
	ID            uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`                              // 自增主键
	DeviceID      int64     `gorm:"unique;column:device_id;type:bigint;not null"`                                    // 设备id
	AppID         uint64    `gorm:"index:idx_app_id_user_id;column:app_id;type:bigint unsigned;not null"`            // app_id
	UserID        uint64    `gorm:"index:idx_app_id_user_id;column:user_id;type:bigint unsigned;not null;default:0"` // 账户id
	Type          int8      `gorm:"column:type;type:tinyint;not null"`                                               // 设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
	Brand         string    `gorm:"column:brand;type:varchar(20);not null"`                                          // 手机厂商
	Model         string    `gorm:"column:model;type:varchar(20);not null"`                                          // 机型
	SystemVersion string    `gorm:"column:system_version;type:varchar(10);not null"`                                 // 系统版本
	SdkVersion    string    `gorm:"column:sdk_version;type:varchar(10);not null"`                                    // app版本
	Status        int8      `gorm:"column:status;type:tinyint;not null;default:0"`                                   // 在线状态，0：离线；1：在线
	ConnAddr      string    `gorm:"column:conn_addr;type:varchar(25);not null"`                                      // 连接层服务器地址
	ConnFd        int64     `gorm:"column:conn_fd;type:bigint;not null"`                                             // TCP连接对应的文件描述符
	CreateTime    time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP"`             // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP"`             // 更新时间
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

// DeviceAck 设备消息同步序列号
type DeviceAck struct {
	ID         uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`                  // 自增主键
	DeviceID   uint64    `gorm:"unique;column:device_id;type:bigint unsigned;not null"`               // 设备id
	Ack        uint64    `gorm:"column:ack;type:bigint unsigned;not null;default:0"`                  // 收到消息确认号
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 更新时间
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

// Gid 分布式自增主键
type Gid struct {
	ID          uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`                  // 自增主键
	BusinessID  string    `gorm:"unique;column:business_id;type:varchar(128);not null"`                // 业务id
	MaxID       uint64    `gorm:"column:max_id;type:bigint unsigned;not null;default:0"`               // 最大id
	Step        uint      `gorm:"column:step;type:int unsigned;not null;default:1000"`                 // 步长
	Description string    `gorm:"column:description;type:varchar(255);not null"`                       // 描述
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 更新时间
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

// Group 群组
type Group struct {
	ID           uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`        // 自增主键
	Name         string    `gorm:"column:name;type:varchar(50);not null;default:''"`          // 群组名称
	Introduction string    `gorm:"column:introduction;type:varchar(255);not null;default:''"` // 群组简介
	Extra        string    `gorm:"column:extra;type:varchar(1024);not null;default:''"`       // 附加属性
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime"`                           // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime"`                           // 更新时间
	DeletedAt    time.Time `gorm:"column:deleted_at;type:datetime"`                           // 删除时间
}

// GroupColumns get sql column name.获取数据库列名
var GroupColumns = struct {
	ID           string
	Name         string
	Introduction string
	Extra        string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "id",
	Name:         "name",
	Introduction: "introduction",
	Extra:        "extra",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// GroupUser 群组成员关系
type GroupUser struct {
	ID           uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`                                             // 自增主键
	GroupID      uint64    `gorm:"uniqueIndex:uk_group_id_user_id;column:group_id;type:bigint unsigned;not null"`                  // 组id
	UserID       uint64    `gorm:"uniqueIndex:uk_group_id_user_id;index:idx_user_id;column:user_id;type:bigint unsigned;not null"` // 用户id
	UserShowname string    `gorm:"column:user_showname;type:varchar(20);not null"`                                                 // 用户在群组的昵称
	Extra        string    `gorm:"column:extra;type:varchar(1024);not null"`                                                       // 附加属性
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime"`                                                                // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime"`                                                                // 更新时间
	DeletedAt    time.Time `gorm:"index:deleted_at;column:deleted_at;type:datetime"`                                               // 删除时间
}

// GroupUserColumns get sql column name.获取数据库列名
var GroupUserColumns = struct {
	ID           string
	GroupID      string
	UserID       string
	UserShowname string
	Extra        string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "id",
	GroupID:      "group_id",
	UserID:       "user_id",
	UserShowname: "user_showname",
	Extra:        "extra",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// Message 消息
type Message struct {
	ID           uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`
	SenderID     int       `gorm:"column:sender_id;type:int;not null;default:0"`             // 发送人ID
	SenderType   int8      `gorm:"column:sender_type;type:tinyint;not null;default:0"`       // 发送人类型
	ReceiverID   int       `gorm:"column:receiver_id;type:int;not null;default:0"`           // 接收人ID，单聊则为user_id,群聊则为group_id
	ReceiverType int8      `gorm:"column:receiver_type;type:tinyint;not null;default:0"`     // 接收人类型
	ToUserIDs    string    `gorm:"column:to_user_ids;type:varchar(255);not null;default:''"` // 需要@的用户，多个用,分割
	MessageType  int       `gorm:"column:message_type;type:int;not null;default:0"`          // 消息类型
	Status       int8      `gorm:"column:status;type:tinyint;not null;default:1"`            // 状态：1-正常，2-撤回，0-删除
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime"`                          // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime"`                          // 更新时间
	DeletedAt    time.Time `gorm:"column:deleted_at;type:datetime"`                          // 删除时间
}

// MessageColumns get sql column name.获取数据库列名
var MessageColumns = struct {
	ID           string
	SenderID     string
	SenderType   string
	ReceiverID   string
	ReceiverType string
	ToUserIDs    string
	MessageType  string
	Status       string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "id",
	SenderID:     "sender_id",
	SenderType:   "sender_type",
	ReceiverID:   "receiver_id",
	ReceiverType: "receiver_type",
	ToUserIDs:    "to_user_ids",
	MessageType:  "message_type",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// User 用户
type User struct {
	ID        uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`       // 自增主键
	Nickname  string    `gorm:"column:nickname;type:varchar(20);not null;default:''"`     // 昵称
	Password  string    `gorm:"column:password;type:varchar(20);not null;default:''"`     // 密码
	Key       string    `gorm:"column:key;type:varchar(20);not null;default:''"`          // 加密key
	Gender    bool      `gorm:"column:gender;type:tinyint(1);not null;default:0"`         // 性别，0:未知；1:男；2:女
	AvatarURL string    `gorm:"column:avatar_url;type:varchar(1024);not null;default:''"` // 用户头像链接
	Extra     string    `gorm:"column:extra;type:varchar(1024);not null;default:''"`      // 附加属性
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`                          // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`                          // 更新时间
	DeletedAt time.Time `gorm:"index:deleted_at;column:deleted_at;type:datetime"`         // 删除时间
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID        string
	Nickname  string
	Password  string
	Key       string
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
	Key:       "key",
	Gender:    "gender",
	AvatarURL: "avatar_url",
	Extra:     "extra",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}
