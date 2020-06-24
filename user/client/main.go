package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/consul/v2"

	proto "shopping/user/proto/user"
)

func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	service := grpc.NewService(
		service.Name("go.micro.service.client"),
		service.Version("latest"),
		service.Registry(consulReg),
	)
	service.Init()

	userService := proto.NewUserService("go.micro.service.user", service.Client())

	user := proto.User{
		Id:       0,
		Name:     "wxmsummer",
		Phone:    "12346",
		Password: "12346",
		Points:   0,
		Level:    0,
	}
	rsp, err := userService.Register(context.Background(), &proto.RegisterReq{User: &user})
	if err != nil {
		fmt.Println("userService.Register err=", err)
	}
	fmt.Println("rsp:", rsp)
}
