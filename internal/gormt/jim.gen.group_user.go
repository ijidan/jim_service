package gormt

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GroupUserMgr struct {
	*_BaseMgr
}

// GroupUserMgr open func
func GroupUserMgr(db *gorm.DB) *_GroupUserMgr {
	if db == nil {
		panic(fmt.Errorf("GroupUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("group_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupUserMgr) GetTableName() string {
	return "group_user"
}

// Reset 重置gorm会话
func (obj *_GroupUserMgr) Reset() *_GroupUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GroupUserMgr) Get() (result GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupUserMgr) Gets() (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GroupUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_GroupUserMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGroupID group_id获取 组id
func (obj *_GroupUserMgr) WithGroupID(groupID uint64) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithUserID user_id获取 用户id
func (obj *_GroupUserMgr) WithUserID(userID uint64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserShowName user_show_name获取 用户在群组的昵称
func (obj *_GroupUserMgr) WithUserShowName(userShowName string) Option {
	return optionFunc(func(o *options) { o.query["user_show_name"] = userShowName })
}

// WithExtra extra获取 附加属性
func (obj *_GroupUserMgr) WithExtra(extra string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_GroupUserMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_GroupUserMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GroupUserMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_GroupUserMgr) GetByOption(opts ...Option) (result GroupUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GroupUserMgr) GetByOptions(opts ...Option) (results []*GroupUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 自增主键
func (obj *_GroupUserMgr) GetFromID(id uint64) (result GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_GroupUserMgr) GetBatchFromID(ids []uint64) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容 组id
func (obj *_GroupUserMgr) GetFromGroupID(groupID uint64) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找 组id
func (obj *_GroupUserMgr) GetBatchFromGroupID(groupIDs []uint64) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_GroupUserMgr) GetFromUserID(userID uint64) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户id
func (obj *_GroupUserMgr) GetBatchFromUserID(userIDs []uint64) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserShowName 通过user_show_name获取内容 用户在群组的昵称
func (obj *_GroupUserMgr) GetFromUserShowName(userShowName string) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`user_show_name` = ?", userShowName).Find(&results).Error

	return
}

// GetBatchFromUserShowName 批量查找 用户在群组的昵称
func (obj *_GroupUserMgr) GetBatchFromUserShowName(userShowNames []string) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`user_show_name` IN (?)", userShowNames).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 附加属性
func (obj *_GroupUserMgr) GetFromExtra(extra string) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 附加属性
func (obj *_GroupUserMgr) GetBatchFromExtra(extras []string) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_GroupUserMgr) GetFromCreatedAt(createdAt time.Time) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GroupUserMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_GroupUserMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_GroupUserMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GroupUserMgr) GetFromDeletedAt(deletedAt time.Time) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GroupUserMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GroupUserMgr) FetchByPrimaryKey(id uint64) (result GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUkGroupIDUserID primary or index 获取唯一内容
func (obj *_GroupUserMgr) FetchUniqueIndexByUkGroupIDUserID(groupID uint64, userID uint64) (result GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`group_id` = ? AND `user_id` = ?", groupID, userID).Find(&result).Error

	return
}

// FetchIndexByIDxUserID  获取多个内容
func (obj *_GroupUserMgr) FetchIndexByIDxUserID(userID uint64) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByDeletedAt  获取多个内容
func (obj *_GroupUserMgr) FetchIndexByDeletedAt(deletedAt time.Time) (results []*GroupUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupUser{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
