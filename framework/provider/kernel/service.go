package kernel

import (
	"net/http"

	"github.com/zhixian0949/five/framework/gin"
)

// 引擎服务
type FiveKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewFiveKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &FiveKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *FiveKernelService) HttpEngine() http.Handler {
	return s.engine
}
