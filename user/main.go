package main

import (
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"shopping/user/handler"
	"shopping/user/model"
	"shopping/user/repository"

	user "shopping/user/proto/user"
)

func main() {


	db,err := CreateConnection()
	if err != nil {
		log.Fatal("connection error : %v \n" , err)
	}
	defer db.Close()

	db.AutoMigrate(&model.User{})

	repo := &repository.User{Db: db}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserServiceHandler(service.Server(), &handler.User{repo})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
