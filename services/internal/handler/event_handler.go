package handler

import (
	"encoding/json"
	"fmt"
	"gokart/internal/model"
	"gokart/internal/producer"
	"gokart/internal/storage"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateEventRequest struct {
	Type    string `json:"type"`
	Source  string `json:"source"`
	Payload map[string]interface{}
}

type EventHandler struct {
	Producer producer.EventProducer
}

func (h *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req CreateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Type == "" || req.Source == "" {
		http.Error(w, "Type and Source are required", http.StatusBadRequest)
		return
	}

	newEvent := model.Event{
		ID:        uuid.New().String(),
		EventId:   uuid.New().String(),
		Type:      req.Type,
		Source:    req.Source,
		Timestamp: time.Now().Unix(),
		Payload:   req.Payload,
	}

	if err := h.Producer.Produce(r.Context(), newEvent); err != nil {
		fmt.Println(err)	
		http.Error(w, "Failed to produce event", http.StatusInternalServerError)
		return
	}

	storage.Events = append(storage.Events, newEvent)

	// Save to JSON file
	if err := storage.SaveToJSON(); err != nil {
		http.Error(w, "Failed to save event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
