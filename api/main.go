package api

import (
	"bubble/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	todoList = &[]model.Bubble{}
	todo     = &model.Bubble{}
)

// 默认页面
func Default(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)

}

// 添加数据
func Add(context *gin.Context) {
	context.BindJSON(todo)
	if err := model.DB.Debug().Create(todo).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

// 删除数据
func Delete(context *gin.Context) {
	id, ok := context.Params.Get("id")
	fmt.Println(id)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
		return
	}
	if err := model.DB.Debug().Where("id=?", id).Delete(todo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

// 编辑数据
func Put(context *gin.Context) {
	// 验证参数的有效性
	id, ok := context.Params.Get("id")
	fmt.Println("获取到的 ID", id)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
		return
	}
	// 根据参数查询数据库中是否存在该条记录
	if err := model.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
	// 获取请求参数并绑定
	context.BindJSON(todo)
	if err := model.DB.Debug().Save(todo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

// 查询数据
func Get(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
		return
	}
	if err := model.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

// 查看数据列表（无搜索条件）
func All(context *gin.Context) {
	if err := model.DB.Debug().Find(todoList).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todoList)
	}
}
