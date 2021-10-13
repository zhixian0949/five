package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// FiveContainer 是服务容器的具体实现
type FiveContainer struct {
	Container // 强制要求 HadeContainer 实现 Container 接口
	// providers 存储注册的服务提供者，key 为字符串凭证
	providers map[string]ServiceProvider
	// instance 存储具体的实例，key 为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

// NewHadeContainer 创建一个服务容器
func NewFiveContainer() *FiveContainer {
	return &FiveContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (five *FiveContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range five.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind 将服务容器和关键字做了绑定
func (five *FiveContainer) Bind(provider ServiceProvider) error {
	five.lock.Lock()
	defer five.lock.Unlock()
	key := provider.Name()

	five.providers[key] = provider

	if provider.IsDefer() == false {
		if err := provider.Boot(five); err != nil {
			return err
		}
		params := provider.Params(five)
		method := provider.Register(five)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		five.instances[key] = instance
	}
	return nil
}

// Make 方式调用内部的 make 实现
func (five *FiveContainer) Make(key string) (interface{}, error) {
	return five.make(key, nil, false)
}

func (five *FiveContainer) MustMake(key string) interface{} {
	serv, err := five.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

// MakeNew 方式使用内部的 make 初始化
func (five *FiveContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return five.make(key, params, true)
}

// 真正的实例化一个服务
func (five *FiveContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	five.lock.RLock()
	defer five.lock.RUnlock()

	sp := five.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return five.newInstance(sp, params)
	}

	if ins, ok := five.instances[key]; ok {
		return ins, nil
	}

	inst, err := five.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	five.instances[key] = inst
	return inst, nil
}

func (five *FiveContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(five); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(five)
	}
	method := sp.Register(five)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

func (five *FiveContainer) findServiceProvider(key string) ServiceProvider {
	five.lock.RLock()
	defer five.lock.RUnlock()
	if sp, ok := five.providers[key]; ok {
		return sp
	}
	return nil
}
