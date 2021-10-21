package http

import (
	"github.com/zhixian0949/five/app/http/module/demo"
	"github.com/zhixian0949/five/framework/gin"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
