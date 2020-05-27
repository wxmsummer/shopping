package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	comment "shopping/comment/proto/comment"
)

type Comment struct{}

func (e *Comment) Handle(ctx context.Context, msg *comment.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *comment.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
