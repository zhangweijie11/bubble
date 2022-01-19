package main

import (
	"bubble/api"
	"bubble/model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 初始化连接
func InitGin() *gin.Engine {
	// 创建基础链接
	r := gin.Default()
	// 静态文件路径
	r.Static("/static", "static")
	// 模版文件所在路径
	r.LoadHTMLGlob("templates/*")
	return r
}

func main() {
	// 初始化数据库
	model.InitPg()
	// 程序关闭退出数据库连接
	defer model.DB.Close()

	// 初始化 gin
	r := InitGin()
	r.GET("/index", api.Default)
	// 建立路由组
	v1Group := r.Group("v1")
	{
		v1Group.GET("/todo", api.All)
		v1Group.POST("/todo", api.Add)
		v1Group.DELETE("/todo/:id", api.Delete)
		v1Group.PUT("/todo/:id", api.Put)
		v1Group.GET("/todo/:id", api.Get)
	}

	runErr := r.Run(":9090")
	if runErr != nil {
		panic(runErr)
	}
}
