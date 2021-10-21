package main

import (
	"github.com/zhixian0949/five/app/console"
	"github.com/zhixian0949/five/app/http"
	"github.com/zhixian0949/five/framework"
	"github.com/zhixian0949/five/framework/provider/app"
	"github.com/zhixian0949/five/framework/provider/kernel"
)

func main() {
	container := framework.NewFiveContainer()
	container.Bind(&app.FiveAppProvider{})

	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.FiveKernelProvider{HttpEngine: engine})
	}
	console.RunCommand(container)
}
