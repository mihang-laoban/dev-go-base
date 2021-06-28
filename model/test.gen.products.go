package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductsMgr struct {
	*_BaseMgr
}

// ProductsMgr open func
func ProductsMgr(db *gorm.DB) *_ProductsMgr {
	if db == nil {
		panic(fmt.Errorf("ProductsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("products"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductsMgr) GetTableName() string {
	return "products"
}

// Get 获取
func (obj *_ProductsMgr) Get() (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductsMgr) Gets() (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_ProductsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_ProductsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_ProductsMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_ProductsMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCode code获取
func (obj *_ProductsMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithPrice price获取
func (obj *_ProductsMgr) WithPrice(price uint64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// GetByOption 功能选项模式获取
func (obj *_ProductsMgr) GetByOption(opts ...Option) (result Products, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductsMgr) GetByOptions(opts ...Option) (results []*Products, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary key
func (obj *_ProductsMgr) GetFromID(id int64) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary key
func (obj *_ProductsMgr) GetBatchFromID(ids []int64) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_ProductsMgr) GetFromCreatedAt(createdAt time.Time) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" = ?", createdAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量查找 created time
func (obj *_ProductsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_ProductsMgr) GetFromUpdatedAt(updatedAt time.Time) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" = ?", updatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 updated at
func (obj *_ProductsMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_ProductsMgr) GetFromDeletedAt(deletedAt time.Time) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" = ?", deletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量查找 deleted time
func (obj *_ProductsMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_ProductsMgr) GetFromCode(code string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_ProductsMgr) GetBatchFromCode(codes []string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容
func (obj *_ProductsMgr) GetFromPrice(price uint64) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找
func (obj *_ProductsMgr) GetBatchFromPrice(prices []uint64) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductsMgr) FetchByPrimaryKey(id int64) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
