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
		[]string{"localhost:9092"},
	)

	evtHandler := &handler.EventHandler{
		Producer: prod,
	}

	cons := consumer.NewKafkaConsumer(
		[]string{"localhost:9092"},
		"gokart-group",
		"OrderCreated",
	)

	err := cons.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	

	mux := routes.SetupRoutes(evtHandler)

	fmt.Println("✅ Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

