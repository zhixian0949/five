package app

import "github.com/zhixian0949/five/framework"

// HadeApp 代表 hade 框架的 App 实现
type FiveApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

func (f FiveApp) Version() string {
	return "0.0.3"
}
