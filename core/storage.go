package core

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveWasteToFile(filename string, data []Waste) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func LoadWasteFromFile(filename string) ([]Waste, error) {
	var rawData []Waste

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&rawData); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return CreateManyData(rawData), nil
}
