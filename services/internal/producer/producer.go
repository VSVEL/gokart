package producer

import (
	"context"
	"gokart/internal/model"
)

type EventProducer interface {
	Produce(ctx context.Context, event model.Event) error
}
