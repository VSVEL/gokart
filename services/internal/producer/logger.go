package producer

import (
	"fmt"
	"gokart/internal/model"
)

type StdoutProducer struct{}

func (p *StdoutProducer) Produce(event model.Event) error {
	fmt.Println("Produce event")
	fmt.Println(event)
	return nil
}
