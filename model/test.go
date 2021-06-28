package model

import (
	"gorm.io/gorm"
	"time"
)

// Activities [...]
type Activities struct {
	ID          int       `gorm:"primaryKey;column:Id;type:int(11);not null"`
	Date        time.Time `gorm:"column:Date;type:datetime;not null"`
	Category    string    `gorm:"column:Category;type:text"`
	Description string    `gorm:"column:Description;type:text"`
	Title       string    `gorm:"column:Title;type:text"`
	City        string    `gorm:"column:City;type:text"`
	Venue       string    `gorm:"column:Venue;type:text"`
}

// ActivitiesColumns get sql column name.获取数据库列名
var ActivitiesColumns = struct {
	ID          string
	Date        string
	Category    string
	Description string
	Title       string
	City        string
	Venue       string
}{
	ID:          "Id",
	Date:        "Date",
	Category:    "Category",
	Description: "Description",
	Title:       "Title",
	City:        "City",
	Venue:       "Venue",
}

// Employee [...]
type Employee struct {
	ID    int    `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name  string `gorm:"index:idx_name;column:name;type:varchar(40)"`
	Money int    `gorm:"column:money;type:int(11)"`
}

// EmployeeColumns get sql column name.获取数据库列名
var EmployeeColumns = struct {
	ID    string
	Name  string
	Money string
}{
	ID:    "id",
	Name:  "name",
	Money: "money",
}

// Products [...]
type Products struct {
	gorm.Model
	Code  string `gorm:"column:code;type:varchar(256)"`
	Price uint64 `gorm:"column:price;type:bigint(20) unsigned"`
}

// ProductsColumns get sql column name.获取数据库列名
var ProductsColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Code      string
	Price     string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Code:      "code",
	Price:     "price",
}

// Record [...]
type Record struct {
	ID                 int    `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name               string `gorm:"column:name;type:varchar(50);not null;default:''"`            // 姓名
	Birthday           int    `gorm:"column:birthday;type:int(11)"`                                // 生日
	Phone              int    `gorm:"column:phone;type:int(11)"`                                   // 手机号码
	Distributor        string `gorm:"column:distributor;type:varchar(50);not null;default:''"`     // 所选经销商
	IsBenzOwner        int    `gorm:"column:is_benz_owner;type:int(11);not null;default:0"`        // 是否奔驰车主
	CurrentCarType     string `gorm:"column:current_car_type;type:varchar(50);not null;default:-"` // 现有车型
	DesiredCarType     string `gorm:"column:desired_car_type;type:varchar(50);not null;default:-"` // 意向车型
	NeedTrial          int    `gorm:"column:need_trial;type:int(11);not null;default:0"`           // 是否需要试驾
	WithGirlfriend     int    `gorm:"column:with_girlfriend;type:int(11);not null;default:0"`      // 是否闺蜜同行
	GirlfriendName     string `gorm:"column:girlfriend_name;type:varchar(50);not null;default:''"` // 闺蜜姓名
	GirlfriendBirthday int    `gorm:"column:girlfriend_birthday;type:int(11)"`                     // 闺蜜生日
	GirlfriendPhone    int    `gorm:"column:girlfriend_phone;type:int(11)"`                        // 闺蜜手机号
	RegisterTime       int    `gorm:"column:register_time;type:int(11);not null;default:0"`        // 注册时间
}

// RecordColumns get sql column name.获取数据库列名
var RecordColumns = struct {
	ID                 string
	Name               string
	Birthday           string
	Phone              string
	Distributor        string
	IsBenzOwner        string
	CurrentCarType     string
	DesiredCarType     string
	NeedTrial          string
	WithGirlfriend     string
	GirlfriendName     string
	GirlfriendBirthday string
	GirlfriendPhone    string
	RegisterTime       string
}{
	ID:                 "id",
	Name:               "name",
	Birthday:           "birthday",
	Phone:              "phone",
	Distributor:        "distributor",
	IsBenzOwner:        "is_benz_owner",
	CurrentCarType:     "current_car_type",
	DesiredCarType:     "desired_car_type",
	NeedTrial:          "need_trial",
	WithGirlfriend:     "with_girlfriend",
	GirlfriendName:     "girlfriend_name",
	GirlfriendBirthday: "girlfriend_birthday",
	GirlfriendPhone:    "girlfriend_phone",
	RegisterTime:       "register_time",
}

// Test [...]
type Test struct {
	ID int `gorm:"column:id;type:int(11)"`
}

// TestColumns get sql column name.获取数据库列名
var TestColumns = struct {
	ID string
}{
	ID: "id",
}
