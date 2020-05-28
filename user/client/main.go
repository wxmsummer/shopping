package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"

	proto "shopping/user/proto/user"
)

func main() {

	consulReg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.service.client"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		)
	service.Init()

	userService := proto.NewUserService("go.micro.service.user", service.Client())

	user := proto.User{
		Id:       0,
		Name:     "wxmsummer",
		Phone:    "12345",
		Password: "12345",
		Points:   0,
		Level:    0,
	}
	rsp, err := userService.Register(context.Background(), &proto.RegisterReq{User: &user})
	if err != nil {
		fmt.Println("userService.Register err=",err)
	}
	fmt.Println("rsp:", rsp)
}