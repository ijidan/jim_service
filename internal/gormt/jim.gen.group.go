package gormt

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GroupMgr struct {
	*_BaseMgr
}

// GroupMgr open func
func GroupMgr(db *gorm.DB) *_GroupMgr {
	if db == nil {
		panic(fmt.Errorf("GroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupMgr) GetTableName() string {
	return "group"
}

// Reset 重置gorm会话
func (obj *_GroupMgr) Reset() *_GroupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GroupMgr) Get() (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupMgr) Gets() (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GroupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Group{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_GroupMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 群组名称
func (obj *_GroupMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIntroduction introduction获取 群组简介
func (obj *_GroupMgr) WithIntroduction(introduction string) Option {
	return optionFunc(func(o *options) { o.query["introduction"] = introduction })
}

// WithExtra extra获取 附加属性
func (obj *_GroupMgr) WithExtra(extra string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithUserID user_id获取 群主ID
func (obj *_GroupMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_GroupMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_GroupMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GroupMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_GroupMgr) GetByOption(opts ...Option) (result Group, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GroupMgr) GetByOptions(opts ...Option) (results []*Group, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 自增主键
func (obj *_GroupMgr) GetFromID(id uint64) (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_GroupMgr) GetBatchFromID(ids []uint64) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 群组名称
func (obj *_GroupMgr) GetFromName(name string) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 群组名称
func (obj *_GroupMgr) GetBatchFromName(names []string) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromIntroduction 通过introduction获取内容 群组简介
func (obj *_GroupMgr) GetFromIntroduction(introduction string) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`introduction` = ?", introduction).Find(&results).Error

	return
}

// GetBatchFromIntroduction 批量查找 群组简介
func (obj *_GroupMgr) GetBatchFromIntroduction(introductions []string) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`introduction` IN (?)", introductions).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 附加属性
func (obj *_GroupMgr) GetFromExtra(extra string) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 附加属性
func (obj *_GroupMgr) GetBatchFromExtra(extras []string) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 群主ID
func (obj *_GroupMgr) GetFromUserID(userID int64) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 群主ID
func (obj *_GroupMgr) GetBatchFromUserID(userIDs []int64) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_GroupMgr) GetFromCreatedAt(createdAt time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GroupMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_GroupMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_GroupMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GroupMgr) GetFromDeletedAt(deletedAt time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GroupMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GroupMgr) FetchByPrimaryKey(id uint64) (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Group{}).Where("`id` = ?", id).Find(&result).Error

	return
}
