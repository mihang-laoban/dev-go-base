package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ActivitiesMgr struct {
	*_BaseMgr
}

// ActivitiesMgr open func
func ActivitiesMgr(db *gorm.DB) *_ActivitiesMgr {
	if db == nil {
		panic(fmt.Errorf("ActivitiesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ActivitiesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Activities"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ActivitiesMgr) GetTableName() string {
	return "Activities"
}

// Get 获取
func (obj *_ActivitiesMgr) Get() (result Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ActivitiesMgr) Gets() (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID Id获取
func (obj *_ActivitiesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["Id"] = id })
}

// WithDate Date获取
func (obj *_ActivitiesMgr) WithDate(date time.Time) Option {
	return optionFunc(func(o *options) { o.query["Date"] = date })
}

// WithCategory Category获取
func (obj *_ActivitiesMgr) WithCategory(category string) Option {
	return optionFunc(func(o *options) { o.query["Category"] = category })
}

// WithDescription Description获取
func (obj *_ActivitiesMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["Description"] = description })
}

// WithTitle Title获取
func (obj *_ActivitiesMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["Title"] = title })
}

// WithCity City获取
func (obj *_ActivitiesMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["City"] = city })
}

// WithVenue Venue获取
func (obj *_ActivitiesMgr) WithVenue(venue string) Option {
	return optionFunc(func(o *options) { o.query["Venue"] = venue })
}

// GetByOption 功能选项模式获取
func (obj *_ActivitiesMgr) GetByOption(opts ...Option) (result Activities, err error) {
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
func (obj *_ActivitiesMgr) GetByOptions(opts ...Option) (results []*Activities, err error) {
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

// GetFromID 通过Id获取内容
func (obj *_ActivitiesMgr) GetFromID(id int) (result Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ActivitiesMgr) GetBatchFromID(ids []int) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDate 通过Date获取内容
func (obj *_ActivitiesMgr) GetFromDate(date time.Time) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_ActivitiesMgr) GetBatchFromDate(dates []time.Time) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Date` IN (?)", dates).Find(&results).Error

	return
}

// GetFromCategory 通过Category获取内容
func (obj *_ActivitiesMgr) GetFromCategory(category string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Category` = ?", category).Find(&results).Error

	return
}

// GetBatchFromCategory 批量查找
func (obj *_ActivitiesMgr) GetBatchFromCategory(categorys []string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Category` IN (?)", categorys).Find(&results).Error

	return
}

// GetFromDescription 通过Description获取内容
func (obj *_ActivitiesMgr) GetFromDescription(description string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找
func (obj *_ActivitiesMgr) GetBatchFromDescription(descriptions []string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromTitle 通过Title获取内容
func (obj *_ActivitiesMgr) GetFromTitle(title string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找
func (obj *_ActivitiesMgr) GetBatchFromTitle(titles []string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromCity 通过City获取内容
func (obj *_ActivitiesMgr) GetFromCity(city string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`City` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找
func (obj *_ActivitiesMgr) GetBatchFromCity(citys []string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`City` IN (?)", citys).Find(&results).Error

	return
}

// GetFromVenue 通过Venue获取内容
func (obj *_ActivitiesMgr) GetFromVenue(venue string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Venue` = ?", venue).Find(&results).Error

	return
}

// GetBatchFromVenue 批量查找
func (obj *_ActivitiesMgr) GetBatchFromVenue(venues []string) (results []*Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Venue` IN (?)", venues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ActivitiesMgr) FetchByPrimaryKey(id int) (result Activities, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`Id` = ?", id).Find(&result).Error

	return
}
