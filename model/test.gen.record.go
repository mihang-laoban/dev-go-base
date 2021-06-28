package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RecordMgr struct {
	*_BaseMgr
}

// RecordMgr open func
func RecordMgr(db *gorm.DB) *_RecordMgr {
	if db == nil {
		panic(fmt.Errorf("RecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RecordMgr) GetTableName() string {
	return "record"
}

// Get 获取
func (obj *_RecordMgr) Get() (result Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RecordMgr) Gets() (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_RecordMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 姓名
func (obj *_RecordMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithBirthday birthday获取 生日
func (obj *_RecordMgr) WithBirthday(birthday int) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithPhone phone获取 手机号码
func (obj *_RecordMgr) WithPhone(phone int) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithDistributor distributor获取 所选经销商
func (obj *_RecordMgr) WithDistributor(distributor string) Option {
	return optionFunc(func(o *options) { o.query["distributor"] = distributor })
}

// WithIsBenzOwner is_benz_owner获取 是否奔驰车主
func (obj *_RecordMgr) WithIsBenzOwner(isBenzOwner int) Option {
	return optionFunc(func(o *options) { o.query["is_benz_owner"] = isBenzOwner })
}

// WithCurrentCarType current_car_type获取 现有车型
func (obj *_RecordMgr) WithCurrentCarType(currentCarType string) Option {
	return optionFunc(func(o *options) { o.query["current_car_type"] = currentCarType })
}

// WithDesiredCarType desired_car_type获取 意向车型
func (obj *_RecordMgr) WithDesiredCarType(desiredCarType string) Option {
	return optionFunc(func(o *options) { o.query["desired_car_type"] = desiredCarType })
}

// WithNeedTrial need_trial获取 是否需要试驾
func (obj *_RecordMgr) WithNeedTrial(needTrial int) Option {
	return optionFunc(func(o *options) { o.query["need_trial"] = needTrial })
}

// WithWithGirlfriend with_girlfriend获取 是否闺蜜同行
func (obj *_RecordMgr) WithWithGirlfriend(withGirlfriend int) Option {
	return optionFunc(func(o *options) { o.query["with_girlfriend"] = withGirlfriend })
}

// WithGirlfriendName girlfriend_name获取 闺蜜姓名
func (obj *_RecordMgr) WithGirlfriendName(girlfriendName string) Option {
	return optionFunc(func(o *options) { o.query["girlfriend_name"] = girlfriendName })
}

// WithGirlfriendBirthday girlfriend_birthday获取 闺蜜生日
func (obj *_RecordMgr) WithGirlfriendBirthday(girlfriendBirthday int) Option {
	return optionFunc(func(o *options) { o.query["girlfriend_birthday"] = girlfriendBirthday })
}

// WithGirlfriendPhone girlfriend_phone获取 闺蜜手机号
func (obj *_RecordMgr) WithGirlfriendPhone(girlfriendPhone int) Option {
	return optionFunc(func(o *options) { o.query["girlfriend_phone"] = girlfriendPhone })
}

// WithRegisterTime register_time获取 注册时间
func (obj *_RecordMgr) WithRegisterTime(registerTime int) Option {
	return optionFunc(func(o *options) { o.query["register_time"] = registerTime })
}

// GetByOption 功能选项模式获取
func (obj *_RecordMgr) GetByOption(opts ...Option) (result Record, err error) {
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
func (obj *_RecordMgr) GetByOptions(opts ...Option) (results []*Record, err error) {
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
func (obj *_RecordMgr) GetFromID(id int) (result Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_RecordMgr) GetBatchFromID(ids []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_RecordMgr) GetFromName(name string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_RecordMgr) GetBatchFromName(names []string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromBirthday 通过birthday获取内容 生日
func (obj *_RecordMgr) GetFromBirthday(birthday int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`birthday` = ?", birthday).Find(&results).Error

	return
}

// GetBatchFromBirthday 批量查找 生日
func (obj *_RecordMgr) GetBatchFromBirthday(birthdays []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`birthday` IN (?)", birthdays).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机号码
func (obj *_RecordMgr) GetFromPhone(phone int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 手机号码
func (obj *_RecordMgr) GetBatchFromPhone(phones []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromDistributor 通过distributor获取内容 所选经销商
func (obj *_RecordMgr) GetFromDistributor(distributor string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`distributor` = ?", distributor).Find(&results).Error

	return
}

// GetBatchFromDistributor 批量查找 所选经销商
func (obj *_RecordMgr) GetBatchFromDistributor(distributors []string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`distributor` IN (?)", distributors).Find(&results).Error

	return
}

// GetFromIsBenzOwner 通过is_benz_owner获取内容 是否奔驰车主
func (obj *_RecordMgr) GetFromIsBenzOwner(isBenzOwner int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`is_benz_owner` = ?", isBenzOwner).Find(&results).Error

	return
}

// GetBatchFromIsBenzOwner 批量查找 是否奔驰车主
func (obj *_RecordMgr) GetBatchFromIsBenzOwner(isBenzOwners []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`is_benz_owner` IN (?)", isBenzOwners).Find(&results).Error

	return
}

// GetFromCurrentCarType 通过current_car_type获取内容 现有车型
func (obj *_RecordMgr) GetFromCurrentCarType(currentCarType string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`current_car_type` = ?", currentCarType).Find(&results).Error

	return
}

// GetBatchFromCurrentCarType 批量查找 现有车型
func (obj *_RecordMgr) GetBatchFromCurrentCarType(currentCarTypes []string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`current_car_type` IN (?)", currentCarTypes).Find(&results).Error

	return
}

// GetFromDesiredCarType 通过desired_car_type获取内容 意向车型
func (obj *_RecordMgr) GetFromDesiredCarType(desiredCarType string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`desired_car_type` = ?", desiredCarType).Find(&results).Error

	return
}

// GetBatchFromDesiredCarType 批量查找 意向车型
func (obj *_RecordMgr) GetBatchFromDesiredCarType(desiredCarTypes []string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`desired_car_type` IN (?)", desiredCarTypes).Find(&results).Error

	return
}

// GetFromNeedTrial 通过need_trial获取内容 是否需要试驾
func (obj *_RecordMgr) GetFromNeedTrial(needTrial int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`need_trial` = ?", needTrial).Find(&results).Error

	return
}

// GetBatchFromNeedTrial 批量查找 是否需要试驾
func (obj *_RecordMgr) GetBatchFromNeedTrial(needTrials []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`need_trial` IN (?)", needTrials).Find(&results).Error

	return
}

// GetFromWithGirlfriend 通过with_girlfriend获取内容 是否闺蜜同行
func (obj *_RecordMgr) GetFromWithGirlfriend(withGirlfriend int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`with_girlfriend` = ?", withGirlfriend).Find(&results).Error

	return
}

// GetBatchFromWithGirlfriend 批量查找 是否闺蜜同行
func (obj *_RecordMgr) GetBatchFromWithGirlfriend(withGirlfriends []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`with_girlfriend` IN (?)", withGirlfriends).Find(&results).Error

	return
}

// GetFromGirlfriendName 通过girlfriend_name获取内容 闺蜜姓名
func (obj *_RecordMgr) GetFromGirlfriendName(girlfriendName string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`girlfriend_name` = ?", girlfriendName).Find(&results).Error

	return
}

// GetBatchFromGirlfriendName 批量查找 闺蜜姓名
func (obj *_RecordMgr) GetBatchFromGirlfriendName(girlfriendNames []string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`girlfriend_name` IN (?)", girlfriendNames).Find(&results).Error

	return
}

// GetFromGirlfriendBirthday 通过girlfriend_birthday获取内容 闺蜜生日
func (obj *_RecordMgr) GetFromGirlfriendBirthday(girlfriendBirthday int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`girlfriend_birthday` = ?", girlfriendBirthday).Find(&results).Error

	return
}

// GetBatchFromGirlfriendBirthday 批量查找 闺蜜生日
func (obj *_RecordMgr) GetBatchFromGirlfriendBirthday(girlfriendBirthdays []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`girlfriend_birthday` IN (?)", girlfriendBirthdays).Find(&results).Error

	return
}

// GetFromGirlfriendPhone 通过girlfriend_phone获取内容 闺蜜手机号
func (obj *_RecordMgr) GetFromGirlfriendPhone(girlfriendPhone int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`girlfriend_phone` = ?", girlfriendPhone).Find(&results).Error

	return
}

// GetBatchFromGirlfriendPhone 批量查找 闺蜜手机号
func (obj *_RecordMgr) GetBatchFromGirlfriendPhone(girlfriendPhones []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`girlfriend_phone` IN (?)", girlfriendPhones).Find(&results).Error

	return
}

// GetFromRegisterTime 通过register_time获取内容 注册时间
func (obj *_RecordMgr) GetFromRegisterTime(registerTime int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`register_time` = ?", registerTime).Find(&results).Error

	return
}

// GetBatchFromRegisterTime 批量查找 注册时间
func (obj *_RecordMgr) GetBatchFromRegisterTime(registerTimes []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`register_time` IN (?)", registerTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_RecordMgr) FetchByPrimaryKey(id int) (result Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
