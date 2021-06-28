package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TestMgr struct {
	*_BaseMgr
}

// TestMgr open func
func TestMgr(db *gorm.DB) *_TestMgr {
	if db == nil {
		panic(fmt.Errorf("TestMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TestMgr{_BaseMgr: &_BaseMgr{DB: db.Table("test"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TestMgr) GetTableName() string {
	return "test"
}

// Get 获取
func (obj *_TestMgr) Get() (result Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TestMgr) Gets() (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TestMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// GetByOption 功能选项模式获取
func (obj *_TestMgr) GetByOption(opts ...Option) (result Test, err error) {
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
func (obj *_TestMgr) GetByOptions(opts ...Option) (results []*Test, err error) {
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

// GetFromID 通过id获取内容
func (obj *_TestMgr) GetFromID(id int) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TestMgr) GetBatchFromID(ids []int) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
