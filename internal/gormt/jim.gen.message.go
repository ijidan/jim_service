package gormt

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MessageMgr struct {
	*_BaseMgr
}

// MessageMgr open func
func MessageMgr(db *gorm.DB) *_MessageMgr {
	if db == nil {
		panic(fmt.Errorf("MessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MessageMgr) GetTableName() string {
	return "message"
}

// Reset 重置gorm会话
func (obj *_MessageMgr) Reset() *_MessageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MessageMgr) Get() (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MessageMgr) Gets() (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MessageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Message{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MessageMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSenderID sender_id获取 发送人ID
func (obj *_MessageMgr) WithSenderID(senderID uint64) Option {
	return optionFunc(func(o *options) { o.query["sender_id"] = senderID })
}

// WithSenderType sender_type获取 发送人类型
func (obj *_MessageMgr) WithSenderType(senderType int8) Option {
	return optionFunc(func(o *options) { o.query["sender_type"] = senderType })
}

// WithReceiverID receiver_id获取 接收人ID，单聊则为user_id,群聊则为group_id
func (obj *_MessageMgr) WithReceiverID(receiverID uint64) Option {
	return optionFunc(func(o *options) { o.query["receiver_id"] = receiverID })
}

// WithReceiverType receiver_type获取 接收人类型
func (obj *_MessageMgr) WithReceiverType(receiverType int8) Option {
	return optionFunc(func(o *options) { o.query["receiver_type"] = receiverType })
}

// WithAtUserID at_user_id获取 需要@的用户，多个用,分割
func (obj *_MessageMgr) WithAtUserID(atUserID string) Option {
	return optionFunc(func(o *options) { o.query["at_user_id"] = atUserID })
}

// WithMessageType message_type获取 消息类型
func (obj *_MessageMgr) WithMessageType(messageType int) Option {
	return optionFunc(func(o *options) { o.query["message_type"] = messageType })
}

// WithStatus status获取 状态：1-正常，2-撤回，0-删除
func (obj *_MessageMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_MessageMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_MessageMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_MessageMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_MessageMgr) GetByOption(opts ...Option) (result Message, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MessageMgr) GetByOptions(opts ...Option) (results []*Message, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_MessageMgr) GetFromID(id uint64) (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MessageMgr) GetBatchFromID(ids []uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSenderID 通过sender_id获取内容 发送人ID
func (obj *_MessageMgr) GetFromSenderID(senderID uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`sender_id` = ?", senderID).Find(&results).Error

	return
}

// GetBatchFromSenderID 批量查找 发送人ID
func (obj *_MessageMgr) GetBatchFromSenderID(senderIDs []uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`sender_id` IN (?)", senderIDs).Find(&results).Error

	return
}

// GetFromSenderType 通过sender_type获取内容 发送人类型
func (obj *_MessageMgr) GetFromSenderType(senderType int8) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`sender_type` = ?", senderType).Find(&results).Error

	return
}

// GetBatchFromSenderType 批量查找 发送人类型
func (obj *_MessageMgr) GetBatchFromSenderType(senderTypes []int8) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`sender_type` IN (?)", senderTypes).Find(&results).Error

	return
}

// GetFromReceiverID 通过receiver_id获取内容 接收人ID，单聊则为user_id,群聊则为group_id
func (obj *_MessageMgr) GetFromReceiverID(receiverID uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`receiver_id` = ?", receiverID).Find(&results).Error

	return
}

// GetBatchFromReceiverID 批量查找 接收人ID，单聊则为user_id,群聊则为group_id
func (obj *_MessageMgr) GetBatchFromReceiverID(receiverIDs []uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`receiver_id` IN (?)", receiverIDs).Find(&results).Error

	return
}

// GetFromReceiverType 通过receiver_type获取内容 接收人类型
func (obj *_MessageMgr) GetFromReceiverType(receiverType int8) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`receiver_type` = ?", receiverType).Find(&results).Error

	return
}

// GetBatchFromReceiverType 批量查找 接收人类型
func (obj *_MessageMgr) GetBatchFromReceiverType(receiverTypes []int8) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`receiver_type` IN (?)", receiverTypes).Find(&results).Error

	return
}

// GetFromAtUserID 通过at_user_id获取内容 需要@的用户，多个用,分割
func (obj *_MessageMgr) GetFromAtUserID(atUserID string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`at_user_id` = ?", atUserID).Find(&results).Error

	return
}

// GetBatchFromAtUserID 批量查找 需要@的用户，多个用,分割
func (obj *_MessageMgr) GetBatchFromAtUserID(atUserIDs []string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`at_user_id` IN (?)", atUserIDs).Find(&results).Error

	return
}

// GetFromMessageType 通过message_type获取内容 消息类型
func (obj *_MessageMgr) GetFromMessageType(messageType int) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`message_type` = ?", messageType).Find(&results).Error

	return
}

// GetBatchFromMessageType 批量查找 消息类型
func (obj *_MessageMgr) GetBatchFromMessageType(messageTypes []int) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`message_type` IN (?)", messageTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态：1-正常，2-撤回，0-删除
func (obj *_MessageMgr) GetFromStatus(status int8) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态：1-正常，2-撤回，0-删除
func (obj *_MessageMgr) GetBatchFromStatus(statuss []int8) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_MessageMgr) GetFromCreatedAt(createdAt time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_MessageMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_MessageMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_MessageMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_MessageMgr) GetFromDeletedAt(deletedAt time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_MessageMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MessageMgr) FetchByPrimaryKey(id uint64) (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Message{}).Where("`id` = ?", id).Find(&result).Error

	return
}
