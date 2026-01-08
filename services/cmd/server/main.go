package main

import (
	"fmt"
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

	mux := routes.SetupRoutes(evtHandler)

	fmt.Println("✅ Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

