package main

import (
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"shopping/comment/handler"
	"shopping/comment/subscriber"

	comment "shopping/comment/proto/comment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.comment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	comment.RegisterCommentServiceHandler(service.Server(), new(handler.Comment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.comment", service.Server(), new(subscriber.Comment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
