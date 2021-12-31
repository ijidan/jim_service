package gormt

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _DeviceMgr struct {
	*_BaseMgr
}

// DeviceMgr open func
func DeviceMgr(db *gorm.DB) *_DeviceMgr {
	if db == nil {
		panic(fmt.Errorf("DeviceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DeviceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("device"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DeviceMgr) GetTableName() string {
	return "device"
}

// Reset 重置gorm会话
func (obj *_DeviceMgr) Reset() *_DeviceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DeviceMgr) Get() (result Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DeviceMgr) Gets() (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DeviceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Device{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_DeviceMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDeviceID device_id获取 设备id
func (obj *_DeviceMgr) WithDeviceID(deviceID int64) Option {
	return optionFunc(func(o *options) { o.query["device_id"] = deviceID })
}

// WithAppID app_id获取 app_id
func (obj *_DeviceMgr) WithAppID(appID uint64) Option {
	return optionFunc(func(o *options) { o.query["app_id"] = appID })
}

// WithUserID user_id获取 账户id
func (obj *_DeviceMgr) WithUserID(userID uint64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithType type获取 设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
func (obj *_DeviceMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithBrand brand获取 手机厂商
func (obj *_DeviceMgr) WithBrand(brand string) Option {
	return optionFunc(func(o *options) { o.query["brand"] = brand })
}

// WithModel model获取 机型
func (obj *_DeviceMgr) WithModel(model string) Option {
	return optionFunc(func(o *options) { o.query["model"] = model })
}

// WithSystemVersion system_version获取 系统版本
func (obj *_DeviceMgr) WithSystemVersion(systemVersion string) Option {
	return optionFunc(func(o *options) { o.query["system_version"] = systemVersion })
}

// WithSdkVersion sdk_version获取 app版本
func (obj *_DeviceMgr) WithSdkVersion(sdkVersion string) Option {
	return optionFunc(func(o *options) { o.query["sdk_version"] = sdkVersion })
}

// WithStatus status获取 在线状态，0：离线；1：在线
func (obj *_DeviceMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithConnAddr conn_addr获取 连接层服务器地址
func (obj *_DeviceMgr) WithConnAddr(connAddr string) Option {
	return optionFunc(func(o *options) { o.query["conn_addr"] = connAddr })
}

// WithConnFd conn_fd获取 TCP连接对应的文件描述符
func (obj *_DeviceMgr) WithConnFd(connFd int64) Option {
	return optionFunc(func(o *options) { o.query["conn_fd"] = connFd })
}

// WithCreateTime create_time获取 创建时间
func (obj *_DeviceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_DeviceMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DeviceMgr) GetByOption(opts ...Option) (result Device, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DeviceMgr) GetByOptions(opts ...Option) (results []*Device, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 自增主键
func (obj *_DeviceMgr) GetFromID(id uint64) (result Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_DeviceMgr) GetBatchFromID(ids []uint64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDeviceID 通过device_id获取内容 设备id
func (obj *_DeviceMgr) GetFromDeviceID(deviceID int64) (result Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`device_id` = ?", deviceID).Find(&result).Error

	return
}

// GetBatchFromDeviceID 批量查找 设备id
func (obj *_DeviceMgr) GetBatchFromDeviceID(deviceIDs []int64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`device_id` IN (?)", deviceIDs).Find(&results).Error

	return
}

// GetFromAppID 通过app_id获取内容 app_id
func (obj *_DeviceMgr) GetFromAppID(appID uint64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`app_id` = ?", appID).Find(&results).Error

	return
}

// GetBatchFromAppID 批量查找 app_id
func (obj *_DeviceMgr) GetBatchFromAppID(appIDs []uint64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`app_id` IN (?)", appIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 账户id
func (obj *_DeviceMgr) GetFromUserID(userID uint64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 账户id
func (obj *_DeviceMgr) GetBatchFromUserID(userIDs []uint64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
func (obj *_DeviceMgr) GetFromType(_type int8) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
func (obj *_DeviceMgr) GetBatchFromType(_types []int8) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromBrand 通过brand获取内容 手机厂商
func (obj *_DeviceMgr) GetFromBrand(brand string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`brand` = ?", brand).Find(&results).Error

	return
}

// GetBatchFromBrand 批量查找 手机厂商
func (obj *_DeviceMgr) GetBatchFromBrand(brands []string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`brand` IN (?)", brands).Find(&results).Error

	return
}

// GetFromModel 通过model获取内容 机型
func (obj *_DeviceMgr) GetFromModel(model string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`model` = ?", model).Find(&results).Error

	return
}

// GetBatchFromModel 批量查找 机型
func (obj *_DeviceMgr) GetBatchFromModel(models []string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`model` IN (?)", models).Find(&results).Error

	return
}

// GetFromSystemVersion 通过system_version获取内容 系统版本
func (obj *_DeviceMgr) GetFromSystemVersion(systemVersion string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`system_version` = ?", systemVersion).Find(&results).Error

	return
}

// GetBatchFromSystemVersion 批量查找 系统版本
func (obj *_DeviceMgr) GetBatchFromSystemVersion(systemVersions []string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`system_version` IN (?)", systemVersions).Find(&results).Error

	return
}

// GetFromSdkVersion 通过sdk_version获取内容 app版本
func (obj *_DeviceMgr) GetFromSdkVersion(sdkVersion string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`sdk_version` = ?", sdkVersion).Find(&results).Error

	return
}

// GetBatchFromSdkVersion 批量查找 app版本
func (obj *_DeviceMgr) GetBatchFromSdkVersion(sdkVersions []string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`sdk_version` IN (?)", sdkVersions).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 在线状态，0：离线；1：在线
func (obj *_DeviceMgr) GetFromStatus(status int8) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 在线状态，0：离线；1：在线
func (obj *_DeviceMgr) GetBatchFromStatus(statuss []int8) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromConnAddr 通过conn_addr获取内容 连接层服务器地址
func (obj *_DeviceMgr) GetFromConnAddr(connAddr string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`conn_addr` = ?", connAddr).Find(&results).Error

	return
}

// GetBatchFromConnAddr 批量查找 连接层服务器地址
func (obj *_DeviceMgr) GetBatchFromConnAddr(connAddrs []string) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`conn_addr` IN (?)", connAddrs).Find(&results).Error

	return
}

// GetFromConnFd 通过conn_fd获取内容 TCP连接对应的文件描述符
func (obj *_DeviceMgr) GetFromConnFd(connFd int64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`conn_fd` = ?", connFd).Find(&results).Error

	return
}

// GetBatchFromConnFd 批量查找 TCP连接对应的文件描述符
func (obj *_DeviceMgr) GetBatchFromConnFd(connFds []int64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`conn_fd` IN (?)", connFds).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_DeviceMgr) GetFromCreateTime(createTime time.Time) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_DeviceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_DeviceMgr) GetFromUpdateTime(updateTime time.Time) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_DeviceMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DeviceMgr) FetchByPrimaryKey(id uint64) (result Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUkDeviceID primary or index 获取唯一内容
func (obj *_DeviceMgr) FetchUniqueByUkDeviceID(deviceID int64) (result Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`device_id` = ?", deviceID).Find(&result).Error

	return
}

// FetchIndexByIDxAppIDUserID  获取多个内容
func (obj *_DeviceMgr) FetchIndexByIDxAppIDUserID(appID uint64, userID uint64) (results []*Device, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Device{}).Where("`app_id` = ? AND `user_id` = ?", appID, userID).Find(&results).Error

	return
}
