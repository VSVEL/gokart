package storage

import (
	"encoding/json"
	"gokart/internal/model"
	"os"
)

var Events []model.Event

func SaveToJSON() error {
	file, err := os.Create("events.json")
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(Events)
}
