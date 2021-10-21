package app

import (
	"github.com/zhixian0949/five/framework"
	"github.com/zhixian0949/five/framework/contract"
)

type FiveAppProvider struct {
	BaseFolder string
}

// Register 注册HadeApp方法
func (f *FiveAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeApp
}

// Boot 启动调用
func (f *FiveAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (f *FiveAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (f *FiveAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, f.BaseFolder}
}

// Name 获取字符串凭证
func (f *FiveAppProvider) Name() string {
	return contract.AppKey
}
