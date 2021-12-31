package gormt

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _DeviceAckMgr struct {
	*_BaseMgr
}

// DeviceAckMgr open func
func DeviceAckMgr(db *gorm.DB) *_DeviceAckMgr {
	if db == nil {
		panic(fmt.Errorf("DeviceAckMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DeviceAckMgr{_BaseMgr: &_BaseMgr{DB: db.Table("device_ack"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DeviceAckMgr) GetTableName() string {
	return "device_ack"
}

// Reset 重置gorm会话
func (obj *_DeviceAckMgr) Reset() *_DeviceAckMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DeviceAckMgr) Get() (result DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DeviceAckMgr) Gets() (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DeviceAckMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_DeviceAckMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDeviceID device_id获取 设备id
func (obj *_DeviceAckMgr) WithDeviceID(deviceID uint64) Option {
	return optionFunc(func(o *options) { o.query["device_id"] = deviceID })
}

// WithAck ack获取 收到消息确认号
func (obj *_DeviceAckMgr) WithAck(ack uint64) Option {
	return optionFunc(func(o *options) { o.query["ack"] = ack })
}

// WithCreateTime create_time获取 创建时间
func (obj *_DeviceAckMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_DeviceAckMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DeviceAckMgr) GetByOption(opts ...Option) (result DeviceAck, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DeviceAckMgr) GetByOptions(opts ...Option) (results []*DeviceAck, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 自增主键
func (obj *_DeviceAckMgr) GetFromID(id uint64) (result DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_DeviceAckMgr) GetBatchFromID(ids []uint64) (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDeviceID 通过device_id获取内容 设备id
func (obj *_DeviceAckMgr) GetFromDeviceID(deviceID uint64) (result DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`device_id` = ?", deviceID).Find(&result).Error

	return
}

// GetBatchFromDeviceID 批量查找 设备id
func (obj *_DeviceAckMgr) GetBatchFromDeviceID(deviceIDs []uint64) (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`device_id` IN (?)", deviceIDs).Find(&results).Error

	return
}

// GetFromAck 通过ack获取内容 收到消息确认号
func (obj *_DeviceAckMgr) GetFromAck(ack uint64) (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`ack` = ?", ack).Find(&results).Error

	return
}

// GetBatchFromAck 批量查找 收到消息确认号
func (obj *_DeviceAckMgr) GetBatchFromAck(acks []uint64) (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`ack` IN (?)", acks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_DeviceAckMgr) GetFromCreateTime(createTime time.Time) (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_DeviceAckMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_DeviceAckMgr) GetFromUpdateTime(updateTime time.Time) (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_DeviceAckMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DeviceAckMgr) FetchByPrimaryKey(id uint64) (result DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUkDeviceID primary or index 获取唯一内容
func (obj *_DeviceAckMgr) FetchUniqueByUkDeviceID(deviceID uint64) (result DeviceAck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DeviceAck{}).Where("`device_id` = ?", deviceID).Find(&result).Error

	return
}
