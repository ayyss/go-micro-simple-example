package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"go-micro.dev/v4/client"
)

type UserWrapper struct {
	client.Client
}

// Call 重写父类call方法
func (uw *UserWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	//熔断器配置名称
	cmdName := req.Service() + "." + req.Endpoint()
	//熔断器配置
	hystrixCommandConfig := hystrix.CommandConfig{
		Timeout:                30000, //超时时间
		RequestVolumeThreshold: 20,    //熔断器请求阈值
		ErrorPercentThreshold:  50,    //错误百分比
		SleepWindow:            5000,  //重新检测开启熔断器时间
	}
	//注册绑定熔断器配置与名称
	hystrix.ConfigureCommand(cmdName, hystrixCommandConfig)
	return hystrix.Do(cmdName, func() error {
		return uw.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		return err
	})
}

func NewUserWrapper(c client.Client) client.Client {
	return &UserWrapper{c}
}
