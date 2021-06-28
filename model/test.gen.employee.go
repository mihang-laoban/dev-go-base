package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _EmployeeMgr struct {
	*_BaseMgr
}

// EmployeeMgr open func
func EmployeeMgr(db *gorm.DB) *_EmployeeMgr {
	if db == nil {
		panic(fmt.Errorf("EmployeeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmployeeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("employee"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EmployeeMgr) GetTableName() string {
	return "employee"
}

// Get 获取
func (obj *_EmployeeMgr) Get() (result Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EmployeeMgr) Gets() (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EmployeeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_EmployeeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithMoney money获取
func (obj *_EmployeeMgr) WithMoney(money int) Option {
	return optionFunc(func(o *options) { o.query["money"] = money })
}

// GetByOption 功能选项模式获取
func (obj *_EmployeeMgr) GetByOption(opts ...Option) (result Employee, err error) {
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
func (obj *_EmployeeMgr) GetByOptions(opts ...Option) (results []*Employee, err error) {
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
func (obj *_EmployeeMgr) GetFromID(id int) (result Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_EmployeeMgr) GetBatchFromID(ids []int) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_EmployeeMgr) GetFromName(name string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_EmployeeMgr) GetBatchFromName(names []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromMoney 通过money获取内容
func (obj *_EmployeeMgr) GetFromMoney(money int) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`money` = ?", money).Find(&results).Error

	return
}

// GetBatchFromMoney 批量查找
func (obj *_EmployeeMgr) GetBatchFromMoney(moneys []int) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`money` IN (?)", moneys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EmployeeMgr) FetchByPrimaryKey(id int) (result Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxName  获取多个内容
func (obj *_EmployeeMgr) FetchIndexByIDxName(name string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}
