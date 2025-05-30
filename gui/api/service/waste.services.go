package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"trash-app/core"
	"trash-app/utils"
)

type TrashServices struct {
	// Mutex untuk mengontrol akses ke file
	mu   sync.Mutex
	file string
}

func NewTrashServices() *TrashServices {
	// FormatingSaveData untuk membuat nama file
	filename := utils.FormatingSaveData()

	return &TrashServices{
		file: filename,
	}
}

func (s *TrashServices) CreateWaste(newWaste core.Waste) error {
	var wasteData []core.Waste

	s.mu.Lock()
	defer s.mu.Unlock()

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
	var allWaste []core.Waste

	s.mu.Lock()
	defer s.mu.Unlock()

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

func (s *TrashServices) GetWasteByDate(date string) ([]core.Waste, error) {
	var dataFolder string = "../data/"
	var allWaste []core.Waste

	s.mu.Lock()
	defer s.mu.Unlock()

	// Nama file berdasarkan input tanggal
	filename := dataFolder + date + ".json"
	filename = utils.EnsureJSONExtension(filename)

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	if err := json.Unmarshal(fileContent, &allWaste); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file %s: %w", filename, err)
	}

	return allWaste, nil
}

func (s *TrashServices) UpdateWasteById(id string, updatedWaste core.Waste) (*core.Waste, error) {
	var dataFolder = "../data/"

	s.mu.Lock()
	defer s.mu.Unlock()

	files, err := os.ReadDir(dataFolder)
	if err != nil {
		return nil, fmt.Errorf("failed to read data directory: %w", err)
	}

	for i := 0; i < len(files); i++ {
		file := files[i]
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(dataFolder, file.Name())
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		var wasteList []core.Waste
		if err := json.Unmarshal(fileContent, &wasteList); err != nil {
			return nil, fmt.Errorf("failed to unmarshal file %s: %w", filePath, err)
		}

		for j := 0; j < len(wasteList); j++ {
			if wasteList[j].ID == id {
				// Update fields hanya jika tidak kosong
				if updatedWaste.WasteType != "" {
					wasteList[j].WasteType = updatedWaste.WasteType
				}
				if updatedWaste.RecyclingMethod != "" {
					wasteList[j].RecyclingMethod = updatedWaste.RecyclingMethod
				}
				if updatedWaste.Quantity != 0 {
					wasteList[j].Quantity = updatedWaste.Quantity
				}
				if updatedWaste.Location != "" {
					wasteList[j].Location = updatedWaste.Location
				}
				if updatedWaste.Status != "" {
					wasteList[j].Status = updatedWaste.Status
				}
				wasteList[j].UpdatedAt = time.Now().Format("02-01-2006 15:04:05")

				// Simpan file
				newContent, _ := json.MarshalIndent(wasteList, "", "  ")
				err := os.WriteFile(filePath, newContent, 0644)
				if err != nil {
					return nil, fmt.Errorf("failed to update file %s: %w", filePath, err)
				}

				return &wasteList[j], nil
			}
		}
	}

	return nil, fmt.Errorf("waste data with ID %s not found", id)
}

func (s *TrashServices) DeleteWasteById(id string) error {
	var dataFolder = "../data/"

	s.mu.Lock()
	defer s.mu.Unlock()

	files, err := os.ReadDir(dataFolder)
	if err != nil {
		return fmt.Errorf("failed to read data directory: %w", err)
	}

	found := false

	for i := 0; i < len(files); i++ {
		file := files[i]
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(dataFolder, file.Name())
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		var wasteList []core.Waste
		if err := json.Unmarshal(fileContent, &wasteList); err != nil {
			return fmt.Errorf("failed to unmarshal file %s: %w", filePath, err)
		}

		var newWasteList []core.Waste
		for j := 0; j < len(wasteList); j++ {
			if wasteList[j].ID != id {
				newWasteList = append(newWasteList, wasteList[j])
			} else {
				found = true
			}
		}

		if found {
			// Update file JSON
			newContent, _ := json.MarshalIndent(newWasteList, "", "  ")
			err := os.WriteFile(filePath, newContent, 0644)
			if err != nil {
				return fmt.Errorf("failed to update file %s: %w", filePath, err)
			}
			return nil
		}
	}

	if !found {
		return fmt.Errorf("waste data with ID %s not found", id)
	}

	return nil
}
