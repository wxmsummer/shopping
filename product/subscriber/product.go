package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	product "shopping/product/proto/product"
)

type Product struct{}

func (e *Product) Handle(ctx context.Context, msg *product.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *product.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
