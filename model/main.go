package model

import (
	"bubble/conf"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitPg(dbConf *conf.DbConf) {
	dsn := "host=" + dbConf.Host + " port=" + dbConf.Port + " user=" + dbConf.Username + " dbname=" + dbConf.DbName + " password=" + dbConf.Password + " sslmode=disable"
	fmt.Println(dsn)
	var openErr error
	DB, openErr = gorm.Open("postgres", dsn)
	if openErr != nil {
		panic(openErr)
	}
	err := DB.DB().Ping()
	if err != nil {
		panic(err)
	}
	// 创建数据模型
	DB.AutoMigrate(&Bubble{})
}
