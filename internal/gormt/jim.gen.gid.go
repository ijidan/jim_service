package gormt

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GidMgr struct {
	*_BaseMgr
}

// GidMgr open func
func GidMgr(db *gorm.DB) *_GidMgr {
	if db == nil {
		panic(fmt.Errorf("GidMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GidMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gid"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GidMgr) GetTableName() string {
	return "gid"
}

// Reset 重置gorm会话
func (obj *_GidMgr) Reset() *_GidMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GidMgr) Get() (result Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GidMgr) Gets() (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GidMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gid{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_GidMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBusinessID business_id获取 业务id
func (obj *_GidMgr) WithBusinessID(businessID string) Option {
	return optionFunc(func(o *options) { o.query["business_id"] = businessID })
}

// WithMaxID max_id获取 最大id
func (obj *_GidMgr) WithMaxID(maxID uint64) Option {
	return optionFunc(func(o *options) { o.query["max_id"] = maxID })
}

// WithStep step获取 步长
func (obj *_GidMgr) WithStep(step uint) Option {
	return optionFunc(func(o *options) { o.query["step"] = step })
}

// WithDescription description获取 描述
func (obj *_GidMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithCreateTime create_time获取 创建时间
func (obj *_GidMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_GidMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_GidMgr) GetByOption(opts ...Option) (result Gid, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GidMgr) GetByOptions(opts ...Option) (results []*Gid, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 自增主键
func (obj *_GidMgr) GetFromID(id uint64) (result Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_GidMgr) GetBatchFromID(ids []uint64) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBusinessID 通过business_id获取内容 业务id
func (obj *_GidMgr) GetFromBusinessID(businessID string) (result Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`business_id` = ?", businessID).Find(&result).Error

	return
}

// GetBatchFromBusinessID 批量查找 业务id
func (obj *_GidMgr) GetBatchFromBusinessID(businessIDs []string) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`business_id` IN (?)", businessIDs).Find(&results).Error

	return
}

// GetFromMaxID 通过max_id获取内容 最大id
func (obj *_GidMgr) GetFromMaxID(maxID uint64) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`max_id` = ?", maxID).Find(&results).Error

	return
}

// GetBatchFromMaxID 批量查找 最大id
func (obj *_GidMgr) GetBatchFromMaxID(maxIDs []uint64) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`max_id` IN (?)", maxIDs).Find(&results).Error

	return
}

// GetFromStep 通过step获取内容 步长
func (obj *_GidMgr) GetFromStep(step uint) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`step` = ?", step).Find(&results).Error

	return
}

// GetBatchFromStep 批量查找 步长
func (obj *_GidMgr) GetBatchFromStep(steps []uint) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`step` IN (?)", steps).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 描述
func (obj *_GidMgr) GetFromDescription(description string) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 描述
func (obj *_GidMgr) GetBatchFromDescription(descriptions []string) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_GidMgr) GetFromCreateTime(createTime time.Time) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_GidMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_GidMgr) GetFromUpdateTime(updateTime time.Time) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_GidMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GidMgr) FetchByPrimaryKey(id uint64) (result Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUkBusinessID primary or index 获取唯一内容
func (obj *_GidMgr) FetchUniqueByUkBusinessID(businessID string) (result Gid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gid{}).Where("`business_id` = ?", businessID).Find(&result).Error

	return
}
