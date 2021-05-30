package test

import (
	"dev-go-base/model"
	"fmt"
	"github.com/xxjwxc/public/mysqldb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func GetGorm(dataSourceName string) *gorm.DB {
	database, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{PrepareStmt: false})
	if err != nil {
		panic(err)
	}
	sqlDB, err := database.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return database.Debug()
}

func TestFuncOption(t *testing.T) {
	orm := mysqldb.OnInitDBOrm("root:123@tcp(docker:3308)/test?charset=utf8&parseTime=True&loc=Local&interpolateParams=True") // 推荐方式
	defer orm.OnDestoryDB()

	accountMgr := model.ActivitiesMgr(orm.DB)
	account, _ := accountMgr.GetByOption(accountMgr.WithID(7), accountMgr.WithCategory("FEZDECKCXS")) // 多case条件获取单个
	fmt.Println(account)
}

func TestFuncGet(t *testing.T) {
	model.OpenRelated() // 打开全局预加载 (外键)

	database := GetGorm("root:123@tcp(docker:3308)/test?charset=utf8&parseTime=True&loc=Local&interpolateParams=True")
	defer func() {
		sqldb, _ := database.DB()
		sqldb.Close()
	}()

	accountMgr := model.ActivitiesMgr(database.Where("category = ?", "FEZDECKCXS"))
	account, _ := accountMgr.Get() // 单条获取
	fmt.Println(account)
}
