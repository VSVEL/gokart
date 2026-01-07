package producer

import (
	"fmt"
	"gokart/internal/model"
)

type KafkaProducer struct {
	// Add Kafka writer/config here
	Brokers []string
	Topic   string
}

func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {
	return &KafkaProducer{
		Brokers: brokers,
		Topic:   topic,
	}, nil
}

func (p *KafkaProducer) Produce(event model.Event) error {
	// TODO: Implement Kafka produce logic
	fmt.Printf("Producing event to Kafka: %+v\n", event)
	return nil
}
