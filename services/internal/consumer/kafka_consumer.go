package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"gokart/internal/model"
	"gokart/internal/handler"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic, group string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:   topic,
			GroupID: group,
		}),
	}
}

func (c *KafkaConsumer) Start(ctx context.Context) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		orderID := string(msg.Key)

		var event model.Event
		json.Unmarshal(msg.Value, &event)

		fmt.Println("Kafka event received:", orderID, event.Type)

		if handler.AlreadyProcessed(orderID) {
			fmt.Println("Duplicate order skipped:", orderID)
			continue
		}

		handler.ProcessOrder(orderID, event)

		handler.MarkProcessed(orderID, event.Type)
	}
}