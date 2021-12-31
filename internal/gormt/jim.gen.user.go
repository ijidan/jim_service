package gormt

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_UserMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNickname nickname获取 昵称
func (obj *_UserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithPassword password获取 密码
func (obj *_UserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithGender gender获取 性别，0:未知；1:男；2:女
func (obj *_UserMgr) WithGender(gender string) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// WithAvatarURL avatar_url获取 用户头像链接
func (obj *_UserMgr) WithAvatarURL(avatarURL string) Option {
	return optionFunc(func(o *options) { o.query["avatar_url"] = avatarURL })
}

// WithExtra extra获取 附加属性
func (obj *_UserMgr) WithExtra(extra string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_UserMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_UserMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_UserMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 自增主键
func (obj *_UserMgr) GetFromID(id uint64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_UserMgr) GetBatchFromID(ids []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 昵称
func (obj *_UserMgr) GetFromNickname(nickname string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找 昵称
func (obj *_UserMgr) GetBatchFromNickname(nicknames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_UserMgr) GetFromPassword(password string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_UserMgr) GetBatchFromPassword(passwords []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromGender 通过gender获取内容 性别，0:未知；1:男；2:女
func (obj *_UserMgr) GetFromGender(gender string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`gender` = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量查找 性别，0:未知；1:男；2:女
func (obj *_UserMgr) GetBatchFromGender(genders []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`gender` IN (?)", genders).Find(&results).Error

	return
}

// GetFromAvatarURL 通过avatar_url获取内容 用户头像链接
func (obj *_UserMgr) GetFromAvatarURL(avatarURL string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`avatar_url` = ?", avatarURL).Find(&results).Error

	return
}

// GetBatchFromAvatarURL 批量查找 用户头像链接
func (obj *_UserMgr) GetBatchFromAvatarURL(avatarURLs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`avatar_url` IN (?)", avatarURLs).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 附加属性
func (obj *_UserMgr) GetFromExtra(extra string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 附加属性
func (obj *_UserMgr) GetBatchFromExtra(extras []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_UserMgr) GetFromCreatedAt(createdAt time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_UserMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_UserMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_UserMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_UserMgr) GetFromDeletedAt(deletedAt time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_UserMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(id uint64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexByDeletedAt  获取多个内容
func (obj *_UserMgr) FetchIndexByDeletedAt(deletedAt time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
