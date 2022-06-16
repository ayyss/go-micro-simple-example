package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	//"github.com/go-micro/plugins/v4/registry/nacos"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"micro_demo_gateway/services"
	"micro_demo_gateway/web_pkg"
	"micro_demo_gateway/wrappers"
)

func main() {

	//连接注册中心
	consulRegistry := consul.NewRegistry(registry.Addrs(":8500"))
	//nacosRegistry := nacos.NewRegistry(registry.Addrs("8848"))

	//用户rpc客户端
	userClient := micro.NewService(
		micro.Name("user.client"),
		micro.WrapClient(wrappers.NewUserWrapper), //设置熔断器
	)

	//用户rpc服务端实例
	userRpcService := services.NewUserService("userService", userClient.Client())

	//产品rpc客户端
	productClient := micro.NewService(
		micro.Name("product.client"),
		micro.WrapClient(wrappers.NewProductWrapper),
	)

	//产品rpc服务端实例
	productRpcService := services.NewProductService("productService", productClient.Client())

	//注册网关服务
	gatewayService := web.NewService(
		web.Name("gatewayService"),
		web.Registry(consulRegistry),
		//web.Registry(nacosRegistry),
		web.Address(":9999"),
		web.Handler(web_pkg.InitRouter(userRpcService, productRpcService)),
	)

	gatewayService.Init()
	if err := gatewayService.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
