package main

import (
	"github.com/zhixian0949/five/framework/gin"
	"github.com/zhixian0949/five/framework/middleware"
)

// 注册路由规则
func registerRouter(core *gin.Engine) {
	// 静态路由+HTTP方法匹配
	core.GET("/user/login", middleware.Test1(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test1())
		// 动态路由
		subjectApi.DELETE("/:id", UserLoginController)
		subjectApi.PUT("/:id", UserLoginController)
		subjectApi.GET("/:id", middleware.Test2(), UserLoginController)
		subjectApi.GET("/list/all", UserLoginController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", UserLoginController)
		}
	}
}
