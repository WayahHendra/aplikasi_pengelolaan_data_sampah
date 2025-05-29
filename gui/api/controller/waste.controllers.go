package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"trash-app/core"
	services "trash-app/gui/api/service"
	logger "trash-app/utils"
)

type TrashControllers struct {
	service *services.TrashServices
}

func NewTrashControllers(service *services.TrashServices) *TrashControllers {
	return &TrashControllers{
		service: service,
	}
}

func (c *TrashControllers) CreateWaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Print(logger.Types["error"], "Method not allowed on /waste/add request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}

	var newTrash core.Waste
	err := json.NewDecoder(r.Body).Decode(&newTrash)
	if err != nil {
		log.Print(logger.Types["warning"], "Invalid JSON format on /waste/add request")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "warning",
			"message": "Invalid JSON format",
		})

		return
	}

	err = c.service.CreateWaste(newTrash)
	if err != nil {
		log.Print(logger.Types["error"], err.Error() + " on /waste/add request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}

	log.Print(logger.Types["success"], "Trash data added successfully on /waste request")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Trash data added successfully",
	})

	w.Header().Set("Content-Type", "application/json")
}

func (c *TrashControllers) GetAllWaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Print(logger.Types["error"], "Method not allowed on /waste/get/all request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}	

	wasteData, err := c.service.GetAllWaste()
	if err != nil {
		log.Print(logger.Types["error"], err.Error() + " on /waste/get/all request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}

	log.Print(logger.Types["success"], "Trash data retrieved successfully on /waste/get/all request")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Trash data retrieved successfully on /waste/get/all request",
		"data":    wasteData,
	}

	json.NewEncoder(w).Encode(response)

	w.Header().Set("Content-Type", "application/json")
}

func (c *TrashControllers) GetWasteByDate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	date := parts[3]
	
	if r.Method != http.MethodGet {
		log.Print(logger.Types["error"], "Method not allowed on /waste/get/"+ date +" request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}	

	if len(parts) < 4 {
		log.Print(logger.Types["warning"], "Invalid format date on /waste/get/"+ date +" request")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "warning",
			"message": "Invalid format date",
		})

		return
	}

	wasteData, err := c.service.GetWasteByDate(date)
	if err != nil {
		log.Print(logger.Types["error"], err.Error() + " on /waste/get/"+ date +" request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})
		
		return
	}

	log.Print(logger.Types["success"], "Trash data retrieved successfully on /waste/get/"+ date +" request")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Trash data retrieved successfully on /waste/get/"+ date +" request",
		"data":    wasteData,
	}

	json.NewEncoder(w).Encode(response)

	w.Header().Set("Content-Type", "application/json")
}

func (c *TrashControllers) UpdateWaste(w http.ResponseWriter, r *http.Request) {}

func (c *TrashControllers) DeleteWaste(w http.ResponseWriter, r *http.Request) {}