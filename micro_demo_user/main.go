package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	_ "go-micro.dev/v4/config"
	"go-micro.dev/v4/registry"
	"micro_demo_user/services"
	"micro_demo_user/services_impl"
)

//type mysql struct {
//	Host     string `json:"host"`
//	UserName string `json:"user_name"`
//	Password string `json:"password"`
//}

func main() {

	//连接到注册中心
	consulRegistry := consul.NewRegistry(registry.Addrs(":8500"))

	//consulSource := consul2.NewSource(
	//	consul2.WithAddress(":8500"),
	//	consul2.WithPrefix(""),
	//	consul2.StripPrefix(true),
	//)

	//conf, err := config.NewConfig()
	//if err != nil {
	//	panic(err.Error())
	//}
	//err = conf.Load(consulSource)
	//if err != nil {
	//	panic(err.Error())
	//}

	//mysqlConfig := &mysql{}
	//
	//err = conf.Get("mysql").Scan(&mysqlConfig)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//fmt.Println(mysqlConfig)

	//注册用户服务
	userMicroService := micro.NewService(
		micro.Name("userService"),      //服务名称
		micro.Address(":8081"),         //服务地址与端口
		micro.Registry(consulRegistry), //绑定注册中心
	)

	//开启接收命令行参数
	userMicroService.Init()

	//注册服务方法
	services.RegisterUserServiceHandler(userMicroService.Server(), new(services_impl.UserService))

	//运行服务
	if err := userMicroService.Run(); err != nil {
		fmt.Println(err.Error())
	}

}
