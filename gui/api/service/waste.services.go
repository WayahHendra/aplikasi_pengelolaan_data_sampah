package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
	"trash-app/core"
)

type TrashServices struct {
	mu   sync.Mutex
	file string
}

func NewTrashServices() *TrashServices {
	dataFolder := "../data"

	if _, err := os.Stat(dataFolder); os.IsNotExist(err) {
		os.Mkdir(dataFolder, 0755)
	}

	today := time.Now()
	filename := fmt.Sprintf("%02d-%02d-%04d.json",
		today.Day(),
		today.Month(),
		today.Year(),
	)

	return &TrashServices{
		file: dataFolder + "/" + filename,
	}
}

func (s *TrashServices) CreateWaste(newWaste core.Waste) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var wasteData []core.Waste

	fileContent, err := os.ReadFile(s.file)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if len(fileContent) > 0 {
		if err := json.Unmarshal(fileContent, &wasteData); err != nil {
			return err
		}
	}

	wasteData = append(wasteData, newWaste)
	err = core.SaveWasteToFile(s.file, wasteData)
	if err != nil {
		return err
	}

	return nil
}
