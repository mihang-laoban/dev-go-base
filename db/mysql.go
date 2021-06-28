package db

import (
	"dev-go-base/conf"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	Cons = make(map[string]*MyDb)
	Sync sync.Once
	//DbConf DbConfig
)

type (
	DbConfig struct {
		Db Db
	}
	Db struct {
		Mysql Mysql
	}
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	MyDb struct {
		*gorm.DB
	}
	Product struct {
		gorm.Model
		Code  string
		Price uint
	}
)

func (db *MyDb) Migrate() {
	// Migrate the schema
	db.AutoMigrate(&Product{})
}

func (db *MyDb) Creation() {
	db.Create(&Product{Code: "D42", Price: 100})
}

func (db *MyDb) Read() {
	var product Product
	db.First(&product, "code = ?", "F42") // find product with code D42
	fmt.Println(product)
}

func (db *MyDb) GetAllTableNames() (res []string) {
	var tbl string
	rows, err := db.Raw("show tables").Rows()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&tbl); err != nil {
			fmt.Println(err)
		}
		res = append(res, tbl)
	}
	return
}

type TblCol struct {
	TblName, ColName, ColType, DataType, Extra string
	CharMaxLen                                 int64
}

func (db *MyDb) Execute(sql string) error {
	return db.Exec(sql).Error
}

func (db *MyDb) GetAllCols(names string) (tblCols []TblCol, err error) {
	rows, err := db.Raw("select TABLE_NAME, COLUMN_NAME, COLUMN_TYPE, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH, EXTRA from information_schema.`COLUMNS` where TABLE_NAME = ? ", names).Rows()
	if err != nil {
		return tblCols, errors.New(fmt.Sprintf("%s", err))
	}
	defer rows.Close()
	for rows.Next() {
		var tblCol TblCol
		if scanErr := rows.Scan(&tblCol.TblName, &tblCol.ColName, &tblCol.ColType, &tblCol.DataType, &tblCol.CharMaxLen, &tblCol.Extra); scanErr != nil {
			err = errors.New(fmt.Sprintf("%s", scanErr))
		}
		tblCols = append(tblCols, tblCol)
	}
	return
}

func (db *MyDb) Update() {
	var product Product
	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Println(product)
}

func (db *MyDb) Deletion() {
	var product Product
	db.Delete(&product, 1)
}

func NewDb(dbName string) (*MyDb, error) {
	if Cons[dbName] == nil {
		return nil, errors.New(fmt.Sprintf("%s", "Database should be initiated first."))
	}
	return Cons[dbName], nil
}

func getConfig(confFileName, postfix, confPath string) DbConfig {
	var DbConf DbConfig
	v := conf.ConfigBuilder(confFileName, postfix, confPath)
	if err := v.Unmarshal(&DbConf); err != nil {
		panic(err)
	}
	return DbConf
}

func initDb() {
	dbConf := getConfig("db", "toml", "conf/")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConf.Db.Mysql.User,
		dbConf.Db.Mysql.Password,
		dbConf.Db.Mysql.Host,
		dbConf.Db.Mysql.Port,
		dbConf.Db.Mysql.Database,
	)
	db := NewDbConn(dsn)

	Cons["r"] = &MyDb{db}
}

func NewDbConn(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if d, err := db.DB(); err != nil {
		d.SetConnMaxLifetime(time.Second * 3)
	}
	return db
}

func NewConFromReq(user, password, host, database string, port int) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		database,
	)

	db := NewDbConn(dsn)

	Cons["s"] = &MyDb{db}
}

func Init() {
	Sync.Do(func() {
		initDb()
	})
}
