package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"go-micro.dev/v4/client"
	"time"
)

type ProductWrapper struct {
	client.Client
}

var productOpts client.CallOption = func(options *client.CallOptions) {
	options.RequestTimeout = time.Second * 30
	options.DialTimeout = time.Second * 30
}

// Call 重写父类call方法
func (pw *ProductWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	ctxNew, _ := context.WithTimeout(ctx, time.Second*30)
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
		return pw.Client.Call(ctxNew, req, rsp, productOpts)
	}, func(err error) error {
		return err
	})
}

func NewProductWrapper(c client.Client) client.Client {
	return &UserWrapper{c}
}
