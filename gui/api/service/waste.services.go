package services

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"fmt"
	"os"
	"sync"
	"time"
	"trash-app/core"
	"trash-app/utils"
)

type TrashServices struct {
	mu   sync.Mutex
	file string
}

func NewTrashServices() *TrashServices {
	// Gunakan FormatingSaveData untuk membuat nama file
	filename := utils.FormatingSaveData()

	return &TrashServices{
		file: filename,
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

	// Generate UUID
	uuidv4, _ := utils.GenerateUUIDv4()

	newWaste.ID = uuidv4
	newWaste.CreatedAt = time.Now().Format("02-01-2006 15:04:05")
	
	wasteData = append(wasteData, newWaste)
	err = core.SaveWasteToFile(s.file, wasteData)

	if err != nil {
		return err
	}

	return nil
}

func (s *TrashServices) GetAllWaste() ([]core.Waste, error) {
	var dataFolder string = "../data/"

    s.mu.Lock()
    defer s.mu.Unlock()

    var allWaste []core.Waste

    files, err := os.ReadDir(dataFolder)
    if err != nil {
        return nil, fmt.Errorf("failed to read data directory: %w", err)
    }

    for i := 0; i < len(files); i++ {
        file := files[i]

        // Skip folder
        if file.IsDir() {
            continue
        }

        fileName := utils.EnsureJSONExtension(file.Name())
        if !strings.HasSuffix(fileName, ".json") {
            continue
        }

        filePath := filepath.Join(dataFolder + fileName)
        fileContent, err := os.ReadFile(filePath)
        if err != nil {
            return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
        }

        var wasteData []core.Waste
        if err := json.Unmarshal(fileContent, &wasteData); err != nil {
            return nil, fmt.Errorf("failed to unmarshal file %s: %w", filePath, err)
        }

        allWaste = append(allWaste, wasteData...)
    }

    return allWaste, nil
}