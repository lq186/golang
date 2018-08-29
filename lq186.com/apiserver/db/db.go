package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lq186/golang/lq186.com/apiserver/config"
)

var DB *gorm.DB

const (
	dialect  = "mysql"
	tablePrefix = "api_server_"
)

func init() {
	var err error
	DB, err = gorm.Open(dialect, dbUrl())
	if err != nil {
		panic(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	// auto create tables
	DB.AutoMigrate(&User{}, &Directory{}, &Article{}, &ArticleDetail{})
}

func dbUrl() string {
	return config.Config().DbUrl
}
