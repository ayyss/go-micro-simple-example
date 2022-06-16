package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"micro_demo_product/services"
	"micro_demo_product/services_impl"
)

func main() {
	consulRegistry := consul.NewRegistry(registry.Addrs(":8500"))

	productMicroService := micro.NewService(
		micro.Name("productService"),   //服务名称
		micro.Address(":8089"),         //服务地址与端口
		micro.Registry(consulRegistry), //绑定注册中心
	)

	productMicroService.Init()

	services.RegisterProductServiceHandler(productMicroService.Server(), new(services_impl.ProductService))

	if err := productMicroService.Run(); err != nil {
		fmt.Println(err.Error())
	}

}
