package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/consul/v2"
	db "shopping/comment/database"
	"shopping/comment/handler"
	"shopping/comment/repository"

	proto "shopping/comment/proto/comment"
)

func main() {

	db, err := db.InitDB()
	if err != nil {
		log.Fatal("connection error : %v \n", err)
	}
	defer db.Close()

	repo := &repository.Comment{Db: db}

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	// New Service
	service := grpc.NewService(
		service.Name("go.micro.service.comment"),
		service.Version("latest"),
		service.Registry(consulReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterCommentServiceHandler(service.Server(), &handler.Comment{Repo: repo})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
