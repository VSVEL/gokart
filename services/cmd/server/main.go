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
)

func main() {
	prod := producer.NewKafkaProducer(
		[]string{"localhost:9093"},
	)

	evtHandler := &handler.EventHandler{
		Producer: prod,
	}

	cons := consumer.NewKafkaConsumer(
		[]string{"localhost:9093"},
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
