package handler

import (
	"fmt"
	"gokart/internal/model"
)

func ProcessOrder(orderID string, event model.Event) {
	switch event.Type {
	case "OrderCreated":
		fmt.Println("Order created:", orderID)
	case "OrderCancelled":
		fmt.Println("Order cancelled:", orderID)
	default:
		fmt.Println("Unknown event type:", event.Type)
	}
}

func AlreadyProcessed(orderID string) bool {
	return false
}

func MarkProcessed(orderID string, eventType string) {
	
}
