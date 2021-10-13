package middleware

import (
	"fmt"

	"github.com/zhixian0949/five/framework/gin"
)

func Test1() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test1")
		c.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test1")

	}
}

func Test2() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test2")
		c.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test2")

	}
}
