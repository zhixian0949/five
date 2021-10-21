package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/zhixian0949/five/framework"
	"github.com/zhixian0949/five/framework/util"
)

// HadeApp 代表 hade 框架的 App 实现
type FiveApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

func (f FiveApp) Version() string {
	return "0.0.3"
}

func (f FiveApp) BaseFolder() string {
	if f.baseFolder != "" {
		return f.baseFolder
	}

	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "baseFolder参数，默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}
	return util.GetExecDirectory()
}

func (f FiveApp) ConfigFolder() string {
	return filepath.Join(f.BaseFolder(), "config")
}
func (f FiveApp) LogFolder() string {
	return filepath.Join(f.BaseFolder(), "log")
}
func (f FiveApp) HttpFolder() string {
	return filepath.Join(f.BaseFolder(), "http")
}
func (f FiveApp) ConsoleFolder() string {
	return filepath.Join(f.BaseFolder(), "console")
}
func (f FiveApp) StorageFolder() string {
	return filepath.Join(f.BaseFolder(), "storage")
}
func (f FiveApp) ProviderFolder() string {
	return filepath.Join(f.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (f FiveApp) MiddlewareFolder() string {
	return filepath.Join(f.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (f FiveApp) CommandFolder() string {
	return filepath.Join(f.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (f FiveApp) RuntimeFolder() string {
	return filepath.Join(f.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (f FiveApp) TestFolder() string {
	return filepath.Join(f.BaseFolder(), "test")
}

// NewHadeApp 初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &FiveApp{baseFolder: baseFolder, container: container}, nil
}
