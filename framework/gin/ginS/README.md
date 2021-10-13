# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/zhixian0949/five/framework/gin"
	"github.com/zhixian0949/five/framework/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
