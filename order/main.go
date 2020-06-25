package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/consul/v2"
	db "shopping/order/database"
	"shopping/order/handler"
	"shopping/order/repository"

	proto "shopping/order/proto/order"
)

func main() {

	db, err := db.InitDB()
	if err != nil {
		log.Fatal("connection error : %v \n", err)
	}
	defer db.Close()

	repo := &repository.Order{Db: db}

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	// New Service
	service := grpc.NewService(
		service.Name("go.micro.service.order"),
		service.Version("latest"),
		service.Registry(consulReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterOrderServiceHandler(service.Server(), &handler.Order{Repo: repo})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
