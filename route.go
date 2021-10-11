package main

import (
	"five/framework"
	"five/framework/middleware"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", middleware.Test1(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test1())
		// 动态路由
		subjectApi.Delete("/:id", UserLoginController)
		subjectApi.Put("/:id", UserLoginController)
		subjectApi.Get("/:id", middleware.Test2(), UserLoginController)
		subjectApi.Get("/list/all", UserLoginController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", UserLoginController)
		}
	}
}
