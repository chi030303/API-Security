package router

import (
	"API-Security/controller"
	"API-Security/utils"
	"github.com/gin-gonic/gin"
)

// 将请求封装到Router函数中
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.RequestCounterMiddleware())
	r.Use(utils.ResponseTimeMiddleware())
	r.Use(utils.LogMiddleware())	
	// 学生信息查询接口
	r.GET("/students", controller.SearchInfo)
	return r
}
