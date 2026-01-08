package producer

import (
	"context"
	"fmt"
	"gokart/internal/model"
	"github.com/segmentio/kafka-go"
	"encoding/json"
)

type KafkaProducer struct {
	Brokers []string
	Topics  map[string]string
	writers map[string]*kafka.Writer
}

func NewKafkaProducer(brokers []string) *KafkaProducer {
	writers := make(map[string]*kafka.Writer)

	for _, topic := range []string{
		"OrderCreated",
		"OrderCancelled",
	} {
		writers[topic] = &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.Hash{},
		}
	}

	return &KafkaProducer{writers: writers}
}


func (p *KafkaProducer) Produce(ctx context.Context, event model.Event) error {
	writer, ok := p.writers[event.Type]
	if !ok {
		return fmt.Errorf("no writer for event type: %s", event.Type)
	}

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(event.EventId), // 🔑 ORDERING + IDEMPOTENCY
		Value: value,
	}

	return writer.WriteMessages(ctx, msg)
}

