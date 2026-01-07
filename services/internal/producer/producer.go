package producer

import "gokart/internal/model"

type EventProducer interface {
	Produce(event model.Event) error
}
