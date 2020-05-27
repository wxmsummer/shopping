package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro"
	"shopping/order/handler"
	"shopping/order/subscriber"

	order "shopping/order/proto/order"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderServiceHandler(service.Server(), new(handler.Order))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.order", service.Server(), new(subscriber.Order))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
