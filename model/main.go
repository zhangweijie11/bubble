package model

import "github.com/jinzhu/gorm"

var DB *gorm.DB

// 初始化数据库
func InitPg() {
	dsn := "host=172.30.0.56 port=5432 user=postgres dbname=bubble password=postgres sslmode=disable"
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
