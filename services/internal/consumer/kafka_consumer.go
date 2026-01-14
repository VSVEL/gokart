package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"gokart/internal/handler"
	"gokart/internal/model"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader     *kafka.Reader
	ConsumerID string
}

func NewKafkaConsumer(brokers []string, topic, group, consumerID string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: group,
		}),
		ConsumerID: consumerID,
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

		fmt.Printf("[%s] Kafka event received: OrderID=%s, EventType=%s\n", c.ConsumerID, orderID, event.Type)

		if handler.AlreadyProcessed(orderID) {
			fmt.Printf("[%s] Duplicate order skipped: %s\n", c.ConsumerID, orderID)
			continue
		}

		fmt.Printf("[%s] Processing order: %s\n", c.ConsumerID, orderID)
		handler.ProcessOrder(orderID, event)

		handler.MarkProcessed(orderID, event.Type)
		fmt.Printf("[%s] Order processed successfully: %s\n", c.ConsumerID, orderID)
	}
}
