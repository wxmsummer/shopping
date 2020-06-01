package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/consul/v2"
	db "shopping/product/database"
	"shopping/product/handler"
	"shopping/product/repository"

	proto "shopping/product/proto/product"
)

func main() {

	db, err := db.InitDB()
	if err != nil {
		log.Fatal("connection error : %v \n", err)
	}
	defer db.Close()

	repo := &repository.Product{Db: db}

	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := grpc.NewService(
		service.Name("go.micro.service.product"),
		service.Version("latest"),
		service.Registry(consulReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterProductServiceHandler(service.Server(), &handler.Product{Repo: repo})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
