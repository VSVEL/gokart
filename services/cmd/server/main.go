package main

import (
	"context"
	"fmt"
	"gokart/internal/consumer"
	"gokart/internal/handler"
	"gokart/internal/producer"
	"gokart/internal/routes"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Parse Kafka brokers from environment variable
	// Expected format: "localhost:9092" or "broker1:9092,broker2:9092"
	kafkaBrokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")

	prod := producer.NewKafkaProducer(kafkaBrokers)

	evtHandler := &handler.EventHandler{
		Producer: prod,
	}

	cons := consumer.NewKafkaConsumer(
		kafkaBrokers,
		"OrderCreated",
		"gokart-group",
		"consumer-1",
	)

	go func() {
		log.Println("Kafka consumer started")
		if err := cons.Start(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	mux := routes.SetupRoutes(evtHandler)

	fmt.Println("✅ Server running on http://localhost:8080")

	log.Println("HTTP server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
