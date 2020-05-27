package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"shopping/product/handler"
	"shopping/product/subscriber"

	product "shopping/product/proto/product"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), new(handler.Product))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.product", service.Server(), new(subscriber.Product))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
