package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"trash-app/core"
	services "trash-app/gui/api/service"
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
		log.Println("Method not allowed on /waste request")
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
		log.Println("Invalid JSON format on /waste request")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid JSON format",
		})

		return
	}

	err = c.service.CreateWaste(newTrash)
	if err != nil {
		log.Println(err.Error() + " on /waste request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}

	log.Println("Trash data added successfully on /waste request")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Trash data added successfully",
	})

	w.Header().Set("Content-Type", "application/json")
}

func (c *TrashControllers) GetAllWaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed on /waste request")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}	

	wasteData, err := c.service.GetAllWaste()
	if err != nil {
		log.Println(err.Error() + " on /waste request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}

	log.Println("Trash data retrieved successfully on /waste request")
	w.WriteHeader(http.StatusOK)

	wasteDataJSON, err := json.Marshal(wasteData)

	if err != nil {
		log.Println(err.Error() + " on /waste request")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Method not allowed",
		})

		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Trash data retrieved successfully on /waste request",
		"data":    string(wasteDataJSON),
	})

	w.Header().Set("Content-Type", "application/json")
}