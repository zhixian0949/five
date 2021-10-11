package main

import (
	"context"
	"five/framework"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(3*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 这里做具体的业务
		time.Sleep(2 * time.Second)
		c.Json(200, "ok")
		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}

func UserLoginController(c *framework.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, UserLoginController")
	return nil
}
