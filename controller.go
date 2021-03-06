package main

import (
	"time"

	"github.com/zhixian0949/five/framework/gin"
	"github.com/zhixian0949/five/provider/demo"
)

// func FooControllerHandler(c *gin.Context) error {
// 	finish := make(chan struct{}, 1)
// 	panicChan := make(chan interface{}, 1)

// 	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(3*time.Second))
// 	defer cancel()

// 	go func() {
// 		defer func() {
// 			if p := recover(); p != nil {
// 				panicChan <- p
// 			}
// 		}()
// 		// 这里做具体的业务
// 		time.Sleep(2 * time.Second)
// 		c.SetStatus(200).Json("ok")
// 		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
// 		finish <- struct{}{}
// 	}()

// 	select {
// 	case p := <-panicChan:
// 		c.WriterMux().Lock()
// 		defer c.WriterMux().Unlock()
// 		log.Println(p)
// 		c.SetStatus(500).Json("panic")
// 	case <-finish:
// 		fmt.Println("finish")
// 	case <-durationCtx.Done():
// 		c.WriterMux().Lock()
// 		defer c.WriterMux().Unlock()
// 		c.SetStatus(500).Json("panic")
// 		c.SetHasTimeout()
// 	}
// 	return nil
// }

type Student struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

func UserLoginController(c *gin.Context) {
	// 打印控制器名字
	demoService := c.MustMake(demo.Key).(demo.Service)
	f := demoService.GetFoo()
	time.Sleep(2 * time.Second)
	c.ISetStatus(200).IJson(f)
}
