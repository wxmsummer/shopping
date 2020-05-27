package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"

	proto "shopping/user/proto/user"
)

func main() {
	service := micro.NewService()
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