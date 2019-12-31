package database

import (
	"fmt"
	"testgame/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)

func init() {
	dbConfig := config.Cfg.Database
	db, err := gorm.Open(dbConfig.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host+":"+dbConfig.Port,
		dbConfig.Name))
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if dbConfig.Prefix != "" {
			return dbConfig.Prefix + "_" + defaultTableName
		} else {
			return defaultTableName
		}
	}
	if err != nil {
		panic(err.Error())
	}

	// 全局禁用表名复数
	//db.SingularTable(true)

	//链接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	Db = db
	ping()
}

func Close() {
	if err := Db.Close(); err != nil {
		panic(err.Error())
	}
}

func ping() {
	if err := Db.DB().Ping(); err != nil {
		panic(err.Error())
	}
}
